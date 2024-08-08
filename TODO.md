# Migrating to a backend template TODOs

## The boring stuff
- [ ] readme docs for [vsc versioning](./internal/vcs/)
- [x] readme: getting started
- [ ] swagger docs

## Clean up & Refactor
- [x] refactor middleware
- [ ] SMTP: re-roll credentials, accidentally saved them in the other repo
- [ ] remove movies migrations
- [ ] migrate to new net/http routing
- [ ] add CORS_TRUSTED_ORIGINS to env file

## New Features
- [ ] remove movie routes
- [ ] add tests for authentication
- [x] setup auto migrations with go-embed - [article](https://oscarforner.com/blog/2023-10-10-go-embed-for-migrations/)
- [x] create flag for auto-migrate
- [ ] generate cert for  SSL + add instructions

## Docker & DB
- [x] create docker containers
- [ ] create test/example data
- [ ] docker: create a new database with permissions
- [ ] figure out how to run commands within docker for [migrations](https://github.com/golang-migrate/migrate) + [staticcheck](https://staticcheck.dev/)

## Authentication and authorization
- [ ] endpoint for adding additional activation tokens (see book) 
- [ ] add option for users to reset password (see book)
