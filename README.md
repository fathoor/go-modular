# Go REST API Boilerplate

[![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/static/v1?style=for-the-badge&message=PostgreSQL&color=4169E1&logo=PostgreSQL&logoColor=FFFFFF&label=)](https://www.postgresql.org/)

This is a boilerplate project for building a Modular REST API in Go, aimed to speed up the process of starting a new project and structuring the app.

***

It uses the following major libraries:
- **fiber** - Web Framework: https://github.com/gofiber/fiber
- **sqlx** - Database Access Layer: https://github.com/jmoiron/sqlx
- **pq** - PostgreSQL Driver: https://github.com/lib/pq
- **go-json** - JSON Serialization: https://github.com/goccy/go-json
- **validator** - Request Validation: https://github.com/go-playground/validator

## Usage

A `.env.example` file is provided in the root directory. Copy it to `.env` and adjust the values to your needs.

```bash
# App
APP_PORT=

# Database
DB_DSN="database://user:password@host:port/dbname?sslmode=&TimeZone="
DB_CONN_OPEN=
DB_CONN_IDLE=
DB_CONN_LIFETIME=
```

To initialize new modules, run the following command:

```bash
go run scripts/module.go <module name>
```

To run the app, execute the following command:

```bash
go run cmd/app/main.go
```
