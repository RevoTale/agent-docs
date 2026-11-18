# Go

Use the [Go library choices](../../awesome/go.md) when selecting packages for a matching capability. Prefer the patterns in [100 Go Mistakes](https://github.com/teivah/100-go-mistakes).

## Linting and formatting

Pin [golangci-lint](https://github.com/golangci/golangci-lint) and `govulncheck` in CI and shared development tools.
Use the latest [configuration format](https://golangci-lint.run/docs/configuration/file/) supported by the pinned linter.
Enable `govet`, `staticcheck`, and `lll` with a maximum line length of 120 characters. Enable `gofmt` under
`formatters.enable` using its default settings.

## Taskfile checks

Run each command from every Go module or use an equivalent scope that covers all project packages. Clarify any requested subset when the scope is ambiguous.

| Task | Required behavior |
| --- | --- |
| `task validate` | Run `go build ./...`, `golangci-lint run ./...`, and `govulncheck ./...`. Report formatting differences without writing files. |
| `task fix` | Apply formatting with `golangci-lint fmt ./...`. |
| `task test` | Run `go test -race -count=1 -shuffle=on ./...` once. Run suites that require services or build tags in dedicated tasks with their prerequisites. |

Both `task validate` and `task test` must pass before merge. Use `t.Parallel()` for independent tests and isolate mutable resources.
For database tests, use a private schema, a transaction that rolls back all operations, or unique records with cleanup limited to the test's own data.

## Keep domain entities independent of I/O

### Why

Saving a record or sending a notification inside an entity method couples a business decision to external services.

### Rule

Keep an entity's state and its business rules together. Let an application service load and save the entity and notify other systems.

```go
func (o *Order) Cancel() error {
	if o.Status == "shipped" {
		return errors.New("shipped orders cannot be cancelled")
	}
	o.Status = "cancelled"
	return nil
}
```

### Result

A test can check whether a shipped order may be cancelled with an `Order` value alone. If it needs a mock mailer or repository to make that decision, move those dependencies to the caller.
