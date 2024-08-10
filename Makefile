BINARY_NAME = txplorer
MAIN_PACKAGE_PATH = ./cmd
QUERIES_PACKAGE_PATH = ./internal/persistance/queries

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

## sql: generate all sql related code
.PHONY: sql
sql:
	docker run --rm -v .:/src -w /src sqlc/sqlc generate -f="./sqlc.yml"
