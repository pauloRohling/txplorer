BINARY_NAME = txplorer
MAIN_PACKAGE_PATH = ./cmd
QUERIES_PACKAGE_PATH = ./internal/persistance/queries
SCHEMA_PACKAGE_PATH = ./internal/persistance/schema
POSTGRES_URL = "postgres://postgres:postgres@localhost:5432/txplorer?sslmode=disable"

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## install: install all dependencies
.PHONY: install
install:
	go mod tidy -e

## up: start the database
.PHONY: up
up:
	docker-compose up -d

## down: stop the database
.PHONY: down
down:
	docker-compose down

.PHONY: migrate-up
## migrate-up: execute all migrations
migrate-up:
	docker run --rm -v .:/migrations --network host migrate/migrate -verbose -path=/migrations/$(SCHEMA_PACKAGE_PATH) -database $(POSTGRES_URL) up

.PHONY: migrate-down
## migrate-down: revert all migrations
migrate-down:
	docker run --rm -v .:/migrations --network host migrate/migrate -verbose -path=/migrations/$(SCHEMA_PACKAGE_PATH) -database $(POSTGRES_URL) down

.PHONY: migration
## migration name=?: create a new migration
migration:
	docker run --rm -v .:/migrations --network host migrate/migrate create -ext=sql -dir=/migrations/$(SCHEMA_PACKAGE_PATH) -seq $(name)

## sql: generate all sql related code
.PHONY: sql
sql:
	docker run --rm -v .:/src -w /src sqlc/sqlc generate -f="./sqlc.yml"

## mock: generate mocks
.PHONY: mock
mock:
	docker run --rm -v .:/src -w /src vektra/mockery --all

## test: run all tests
.PHONY: test
test:
	go test -v -race -failfast -buildvcs ./internal/...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=./tmp/coverage.out ./internal/...
	go tool cover -html=./tmp/coverage.out

## run: run the application
.PHONY: run
run:
	go run -v ./cmd/main.go