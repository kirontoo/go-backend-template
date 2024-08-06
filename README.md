# Go Backend Template

This is a backend template built with Go, based on the Greenlight project that is 
covered by [Alex Edwards'](https://www.alexedwards.net/) book 
["Let's Go Further](https://lets-go-further.alexedwards.net/).

Essentially this is the Greenlight project but modified to become a backend
template for future projects.

## Getting Started

### Pre-requisites
Make sure to have these installed on your machine:
- [Go version 1.22.0 or higher](https://go.dev/doc/install)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Migrate](https://github.com/golang-migrate/migrate)

### Clone the Repository
```bash
git clone https://github.com/kirontoo/go-backend-template.git
```

### Install Dependencies
Change dirctory into the repo and use the following command to install all
dependencies.

```bash
go mod download
````

### Set Environment Variables
Make a copy of `.envrc_example` and `.env_example`.  

The `.envrc` will hold the necessary data to run the server.
**.envrc file**
```
export DB_DSN="postgres://username:password@localhost:5432/project"
export SMTP_HOST="smtp.mailtrap.io"
export SMTP_PORT=24
export SMTP_USERNAME="smtp-username"
export SMTP_PASSWORD="smtp-password"
export SMTP_SENDER="my-project <no-reply@myrpoject.example.net>"
```

The `.env` file is for building and starting docker containers.
```
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=database
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=root
```

### Build And Start Containers

Have docker installed on your machine and run:
```bash
docker compose up --buld
```

### Run the APi
```bash
make run/api
```


## Help Commands
Run `make help` to see all make commands.

Use this command to see a list of all env variables available to the server

```bash
make run/api/help
```