# Migrating to a backend template TODOs

## The boring stuff
- [ ] readme docs for [vsc versioning](./internal/vcs/)
- [x] readme: getting started
- [ ] swagger docs

## Clean up & Refactor
- [x] refactor middleware
- [x] SMTP: re-roll credentials, accidentally saved them in the other repo
- [ ] remove movies migrations
- [ ] migrate to new net/http routing
- [x] add CORS_TRUSTED_ORIGINS to env file

## New Features
- [ ] remove movie routes
- [ ] add tests for authentication
    - check out mockery for mock data?
- [x] setup auto migrations with go-embed - [article](https://oscarforner.com/blog/2023-10-10-go-embed-for-migrations/)
- [x] create flag for auto-migrate
- [ ] generate cert for  SSL + add instructions

## Docker & DB
- [x] create docker containers
- [ ] create test/example data
- [ ] docker: create a new database with permissions
- [ ] figure out how to run commands within docker for [migrations](https://github.com/golang-migrate/migrate) + [staticcheck](https://staticcheck.dev/)

## Authentication and authorization
- [x] endpoint for adding additional activation tokens (see book) 
- [x] add option for users to reset password (see book)


## ???
- update docker for testing environment (based on "env" varaible -> development|staging|production)

## Testing
- [] clean up test for require permissions middleware