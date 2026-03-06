include .env
export

MIGRATIONS_DIR := $(PWD)/migrations/sqlite3

# allow overriding the database URL; default points to a file in the
# current directory so it works both locally and on the CI runner
APP_DB ?= sqlite3://$(CURDIR)/test.db

# strip the scheme and convert to an absolute filesystem path. realpath
# with -m will produce a canonical absolute path even if the file does
# not yet exist.
DBFILE := $(patsubst sqlite3://%,%,$(APP_DB))
DBFILE := $(shell realpath -m $(DBFILE))

# rebuild APP_DB from the normalized file location and export it so
# child processes (including `go test` in subpackage directories) see
# the same absolute URL.
APP_DB := sqlite3://$(DBFILE)
export APP_DB

.PHONY: prepare test

prepare:
	# ensure the sqlite file exists (directory already exists)
	touch $(DBFILE)

migrate:
# installed using: go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate \
		-path=$(MIGRATIONS_DIR) \
		-database "$(APP_DB)" \
		up

# `test` builds the database before running migrations/tests
test: prepare
	@echo "using APP_DB=$(APP_DB)"
	migrate -path=$(MIGRATIONS_DIR) -database "$(APP_DB)" up
	gotestsum --format testdox
