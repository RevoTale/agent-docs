# Overview
This module defines baseline engineering rules for Go repositories and services.

# Project structure
```text
<go-repo-root>/
  AGENTS.md
  .golangci.yml|.golangci.yaml|.golangci.toml
```

# Strict rules
- MUST read enforced utility/library choices from [../../awesome/index.md](../../awesome/index.md), [../../awesome/go.md](../../awesome/go.md), and any matching capability awesome files before introducing, replacing, or removing Go libraries.
- MUST enforce `required` entries from matching awesome files for corresponding Go capabilities.
- MUST use `golangci-lint` as the Go linter: https://github.com/golangci/golangci-lint.
- MUST use the latest [golangci-lint configuration format](https://golangci-lint.run/docs/configuration/file/) supported by the pinned golangci-lint version.
- MUST pin the golangci-lint version in CI and shared development tooling.
- MUST enforce a maximum line length of 120 characters through golangci-lint configuration.
- SHOULD prefer patterns from `100 Go Mistakes and How to Avoid Them`: https://github.com/teivah/100-go-mistakes.
- MUST configure the `lll` linter in golangci-lint with line length set to 120.
- MUST enable `govet` and `staticcheck` in golangci-lint. Enable `gofmt` under `formatters.enable` with `formatters.settings.gofmt.simplify: false` to match plain gofmt.
- MUST require `task validate` and `task test` to pass before merge. Run the following commands from each Go module or use an explicit equivalent scope covering all project packages.
- MUST run `go build ./...`, `golangci-lint run ./...`, and `govulncheck ./...` in `task validate`. Pin govulncheck in CI and shared development tooling. Validation must report formatting differences without applying fixes.
- MUST use `golangci-lint fmt ./...` in `task fix` to apply the same formatter locally.
- MUST run the default suite once with `go test -race -count=1 -shuffle=on ./...` in `task test`. Run additional integration or end-to-end suites that need services or build tags in dedicated tasks with their prerequisites, without rerunning the default suite.
- MUST call `t.Parallel()` in independent tests and isolate mutable test resources. Database tests must use a private schema, a rolled-back transaction covering all operations, or unique records with cleanup limited to that test's data.
- MUST ensure entities depend on data, not on a service bucket.

Example:
```go
type Order struct {
	ID     string
	Status string
}

func (o *Order) Cancel() error {
	if o.Status == "shipped" {
		return errors.New("shipped orders cannot be cancelled")
	}
	o.Status = "cancelled"
	return nil
}
```

Anti-example:
```go
type Order struct {
	ID       string
	Status   string
	repo     OrderRepository
	notifier Mailer
}

func (o *Order) Cancel(ctx context.Context) error {
	o.Status = "cancelled"
	if err := o.repo.Save(ctx, o); err != nil {
		return err
	}
	return o.notifier.SendOrderCancelled(o.ID)
}
```

Keep repositories, mailers, and API clients in application services; keep the entity focused on its own state.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask user to choose test/lint scope when scope is ambiguous (`all packages` vs `subset`).
