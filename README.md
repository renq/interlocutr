# Interlocutr

It's a project for fun. Don't look at it unless you want to waste your time.

[![Go](https://github.com/renq/interlocutr/actions/workflows/go.yml/badge.svg)](https://github.com/renq/interlocutr/actions/workflows/go.yml)

## Code snippets (for me)

Before I forgot:

### Run tests:

```
make test
```

### Run integration tests with real DB

```
make test-integration
```

### Run server with live reload

```
make run
```

### Generate swagger

```
make swagger
```

## Quality tools

```
make install-tools
lefthook run lints
make lint
```

## Coverage

```
go test -v -coverprofile=cover.out -coverpkg=./... ./... && go tool cover -html=cover.out
```

## SQLX tutorial

https://jmoiron.github.io/sqlx/
https://dev.to/jones_charles_ad50858dbc0/sqlx-your-go-to-database-toolkit-for-go-developers-53n8


## Fun stuff

### Code visualization

go-callvis .

### How to do transactions in Go

https://threedots.tech/post/database-transactions-in-go/
