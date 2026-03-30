.PHONY: prepare test

export MIGRATIONS_DIR := $(shell pwd)/migrations/sqlite3
export PROD_DB := sqlite3://$(shell pwd)/data/app.db
export TEST_DB_FILE := sqlite3://$(shell pwd)/data/test.db
export TEST_INTEGRATION_DB_FILE := sqlite3://$(shell pwd)/data/test_integration.db

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
	migrate -path="$(MIGRATIONS_DIR)" -database "$(TEST_DB_FILE)" up
	TEST_DB=$(TEST_DB_FILE) gotestsum --format testdox

test-integration:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(TEST_INTEGRATION_DB_FILE)" up
	TEST_INTEGRATION_DB=$(TEST_INTEGRATION_DB_FILE) gotestsum --format testdox ./internal/tests

swagger:
	swag init --outputTypes=json,yaml .

run:
	migrate -path="$(MIGRATIONS_DIR)" -database "$(PROD_DB)" up
	@APP_DB=$(PROD_DB) air

generate-css:
	cd assets; npm install; npx tailwindcss --config tailwind.config.js -i input.css -o interlocutr.css  --minify

build: generate-css
	go build
