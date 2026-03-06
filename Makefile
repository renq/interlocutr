include .env
export

MIGRATIONS_DIR := $(PWD)/migrations/sqlite3

APP_DB ?= sqlite3://$(CURDIR)/test.db
DBFILE := $(patsubst sqlite3://%,%,$(APP_DB))
DBFILE := $(shell realpath -m $(DBFILE))

APP_DB := sqlite3://$(DBFILE)
export APP_DB

.PHONY: prepare test

migrate:
# installed using: go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate \
		-path=$(MIGRATIONS_DIR) \
		-database "$(APP_DB)" \
		up

install-tools:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install gotest.tools/gotestsum@latest

test:
	@echo "using APP_DB=$(APP_DB)"
	migrate -path=$(MIGRATIONS_DIR) -database "$(APP_DB)" up
	gotestsum --format testdox
