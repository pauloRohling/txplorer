# TxPlorer

TxPlorer is a transactional accounting system that allows users to manage their financial accounts and track their
transactions. This project was created as an experiment about how to build a transactional system using Go and
PostgreSQL, while also applying clean architecture principles. One of the main challenges in dealing with transactions
in a clean architecture is how to handle them across multiple repositories without leaking database specific code to the
domain layer.

In this project, a transaction manager interface was used inside each use case, ensuring that no database specific code
was used to rollback or commit a transaction. This interface is also responsible for sharing the transaction object as a
context value to allow repositories to access it. This approach allows the domain layer to be kept clean and focused on
the business logic, while the persistence layer handles the transactional aspects of the application.

## Features

- Create accounts
- Transfer funds between accounts
- View account balances
- View transaction history

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [mockery](https://github.com/vektra/mockery)
- [sqlc](https://github.com/kyleconroy/sqlc)
- [migrate](https://github.com/golang-migrate/migrate)
- [make](https://www.gnu.org/software/make/)

## API

| Endpoint                      | Description                       | Method |
|-------------------------------|-----------------------------------|--------|
| `/api/v1/accounts`            | Creates a new account             | POST   |
| `/api/v1/accounts/{id}`       | Gets an account                   | GET    |
| `/api/v1/accounts/{id}`       | Deletes an account                | DELETE |
| `/api/v1/operations/transfer` | Transfer funds to another account | POST   |
| `/api/v1/operations/deposit`  | Deposit funds to an account       | POST   |
| `/api/v1/operations/withdraw` | Withdraw funds from an account    | POST   |
| `/api/v1/users`               | Changes the password of a user    | POST   |
| `/api/v1/users/login`         | Generates a new access token      | POST   |

## Environment Variables

| Variable                      | Description                         | Default Value | Required |
|-------------------------------|-------------------------------------|---------------|----------|
| `SERVER_PORT`                 | Port to listen on                   | 8080          | false    |
| `SECURITY_SECRET`             | Secret used to sign JWT tokens      | -             | true     |
| `SECURITY_TOKEN_EXPIRATION`   | Expiration time of JWT tokens       | 24h           | false    |
| `DATABASE_HOST`               | Host of the database                | -             | true     |
| `DATABASE_PORT`               | Port of the database                | -             | true     |
| `DATABASE_NAME`               | Name of the database                | -             | true     |
| `DATABASE_USER`               | User to connect to the database     | -             | true     |
| `DATABASE_PASSWORD`           | Password to connect to the database | -             | true     |
| `DATABASE_SSL`                | Whether to use SSL                  | false         | false    |
| `DATABASE_MAX_OPEN_CONNS`     | Maximum number of open connections  | 100           | false    |
| `DATABASE_MAX_IDLE_CONNS`     | Maximum number of idle connections  | 10            | false    |
| `DATABASE_CONN_MAX_LIFETIME`  | Maximum lifetime of a connection    | 5m            | false    |
| `DATABASE_CONN_MAX_IDLE_TIME` | Maximum idle time of a connection   | 30s           | false    |
