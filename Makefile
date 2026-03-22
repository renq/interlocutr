.PHONY: prepare test

export MIGRATIONS_DIR := $(shell pwd)/migrations/sqlite3
export PROD_DB := sqlite3://$(shell pwd)/data/app.db
export TEST_DB := sqlite3://$(shell pwd)/data/test.db

lint:
	go fmt ./...
	golangci-lint run

migrate:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(PROD_DB)" up

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

test:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(TEST_DB)" up
	@APP_DB=$(TEST_DB) gotestsum --format testdox

swagger:
	swag init --outputTypes=json,yaml .

run:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(PROD_DB)" up
	@APP_DB=$(PROD_DB) air
