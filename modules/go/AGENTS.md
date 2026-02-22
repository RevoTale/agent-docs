# Overview
This module defines baseline engineering rules for Go repositories and services.

# Folder Structure
```text
<go-repo-root>/
  AGENTS.md
  .golangci.yml|.golangci.yaml|.golangci.toml
```

# Core Behaviors & Patterns
- Use `golangci-lint` as the Go linter: https://github.com/golangci/golangci-lint
- Enforce a maximum line length of 120 characters through golangci-lint configuration.
- Prefer patterns and recommendations from `100 Go Mistakes and How to Avoid Them`: https://github.com/teivah/100-go-mistakes

# Conventions
- Configure the `lll` linter in golangci-lint with line length set to 120.
- Run golangci-lint against all Go packages in the repository.

# Working Agreements
- `golangci-lint run` must pass for Go changes before merge.
- `go test ./...` must pass for Go changes before merge.
- Fix and validate the code with `gofmt` before merging: the official, opinionated code formatting tool for the Go programming language