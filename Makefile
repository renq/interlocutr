include .env
export

MIGRATIONS_DIR := $(PWD)/migrations/sqlite


.PHONY: test

migrate:
# installed using: go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate \
		-path=migrations/sqlite3 \
		-database "$(APP_DB)" \
		up

test:
	bash -c 'set -a; . ./.env.test; set +a; \
	migrate -path=migrations/sqlite3 -database "$$APP_DB" up && \
	gotestsum --format testdox'
