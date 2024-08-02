# Migrating to a backend template TODOs

- [ ] refactor middleware
- [ ] SMTP: re-roll credentials, accidentally saved them in the other repo
- [ ] docker: create a new database with permissions
- [ ] remove movies migrations
- [ ] add CORS_TRUSTED_ORIGINS to env file
- [ ] add option for users to reset password (see book)
- [ ] endpoint for adding additional activation tokens (see book) 
- [ ] migrate to new net/http routing
- [ ] remove movie routes
- [ ] create example data?
- [ ] add tests for authentication
- [ ] swagger docs
- [ ] docs on how versioning works
- [ ] figure out how to run commands within docker for [migrations](https://github.com/golang-migrate/migrate) + [staticcheck](https://staticcheck.dev/)
- [ ] setup auto migrations with go-embed - [article](https://oscarforner.com/blog/2023-10-10-go-embed-for-migrations/)
