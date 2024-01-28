# goth stack

## Components
- Go + Fiber
- Tailwind
- HTMX
- SQLite

## Extenal dependencies

### Templ
`templ` is used to generate templates in go format. Install the CLI with:

```bash
    go install github.com/a-h/templ/cmd/templ@latest
```

### air (optional)
`air` is used to hot reload the application. Install the CLI with:
```bash
    go install github.com/cosmtrek/air@latest
```

## How to run
Before running be sure to add all required environment variables (see [env example](.env.example)) 

- serve: `go run main.go serve` or `air serve`
- migrate up: `go run main.go migrate:up` or `air migrate:up`
- migrate down: `go run main.go migrate:down N` or `air migrate:down N` where N is the number of migrations down

## Roadmap

### Done
- Go
- HTMx
- Tailwind
- Templ
- Migrations
- OIdC auth
- JWT cookie auth

### Next
- [ ] SQL Autogeneration with [sqlc](https://github.com/sqlc-dev/sqlc)
