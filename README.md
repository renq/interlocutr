# Interlocutr

It's a project for fun. Don't look at it unless you want to waste your time.

## Code snippets (for me)

Before I forgot:

### Run tests:

```
gotestsum --format testdox --watch
```

### Run server with live reload

```
air
```

### Generate swagger

```
swag init
```

## Quality tools

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/evilmartians/lefthook@latest

lefthook run lints
```
