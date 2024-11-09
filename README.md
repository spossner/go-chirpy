# go-chirpy

## Prerequisits
- Go 1.23+
- a running instance of postgres (docker compose file included)
- goose - golang database migration tool

## Setup

### database
Launch postgres database - e.g. in docker using `docker compose up` within the project directory.
Services configured in `docker-compose.yaml`

Connection string
```
postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable
```

### GO packages
Run `go install` within the project root directory

### Configuration files
Copy the example configuration file `config.example.json` into `~/.config/gator/config.json` and update the database connect string as needed.

If you are running postgres in docker using the docker compose file the connect string above is already included in the example.

### Deploy database schema
With a running database push the database schema into the database using goose.
If not yet installed (`goose -version`), install goose:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Finally run up migrations by executing goose from within `sql/schema` with the correct connect string:
```
goose postgres <connect-string> up
```
Use your connect string - e.g. `postgres://postgres:postgres@localhost:5432/chirpy`

### Generating SQL stubs with SQLC
In addition to goose (for db migration) we are using SQLC to generate typesafe stubs to execute our SQL queries.
If not yet installed (`sqlc version`), install SQLC:
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Within the root directory of the project, run `sqlc generate`. SQLC is configued via `sqlc.yaml` and will store the generated go code in `internal/database`.

### Configuration
Copy the `.env.example` file to `.env`. Update the connect string as needed.

_Note:_ on localhost make sure to disable SSL by adding `?sslmode=disable` to the connect string.