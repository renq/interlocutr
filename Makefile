.PHONY: prepare test

export MIGRATIONS_DIR := $(shell pwd)/migrations/sqlite3
export PROD_DB := sqlite3://$(shell pwd)/data/app.db
export TEST_DB := sqlite3://$(shell pwd)/data/test.db

migrate:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(PROD_DB)" up

install-tools:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install gotest.tools/gotestsum@latest

test:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(TEST_DB)" up
	@APP_DB=$(TEST_DB) gotestsum --format testdox
