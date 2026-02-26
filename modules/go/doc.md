# Overview
This module defines baseline engineering rules for Go repositories and services.

# Project structure
```text
<go-repo-root>/
  AGENTS.md
  .golangci.yml|.golangci.yaml|.golangci.toml
```

# Strict rules
- MUST use `golangci-lint` as the Go linter: https://github.com/golangci/golangci-lint.
- MUST enforce a maximum line length of 120 characters through golangci-lint configuration.
- SHOULD prefer patterns from `100 Go Mistakes and How to Avoid Them`: https://github.com/teivah/100-go-mistakes.
- MUST configure the `lll` linter in golangci-lint with line length set to 120.
- MUST run golangci-lint against all Go packages through Taskfile tasks.
- MUST require `task validate` to run `golangci-lint run` and pass for Go changes before merge.
- MUST require `task test` to pass when defined and include `go test ./...` (or an explicit scoped equivalent).
- MUST require `task fix` to run `gofmt` before merge.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask user to choose test/lint scope when scope is ambiguous (`all packages` vs `subset`).
