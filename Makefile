# Include variables from .envrc
include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api -db-dsn=${DB_DSN}

# run/api/help: list all environment variables available for cmd/api application
.PHONY: run/api/help
run/api/help:
	go run ./cmd/api -help

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...## vendor: tidy and vendor dependencies

.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s -X main.version=${VERSION}' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -X main.version=${VERSION}' -o=./bin/linux_amd64/api ./cmd/api

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #


# TODO: replace 'production_host_ip' with your production server's IP
production_host_ip = '161.35.71.158'

## production/connect: connect to the production server
# TODO: replace 'project_name' with your project name
.PHONY: production/connect
production/connect:
	ssh project_name@${production_host_ip}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/api
production/deploy/api:
	rsync -P ./bin/linux_amd64/api project_name@${production_host_ip}:~
	rsync -rP --delete ./migrations project_name@${production_host_ip}:~
	rsync -P ./remote/production/api.service project_name@${production_host_ip}:~
	rsync -P ./remote/production/Caddyfile project_name@${production_host_ip}:~
	ssh -t project_name@${production_host_ip} '\
	migrate -path ~/migrations -database $$DB_DSN up \
	&& sudo mv ~/api.service /etc/systemd/system/ \
	&& sudo systemctl enable api \
	&& sudo systemctl restart api \
	&& sudo mv ~/Caddyfile /etc/caddy/ \
	&& sudo systemctl reload caddy \
	'
