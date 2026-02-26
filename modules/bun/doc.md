# Overview
This module defines baseline engineering rules for Bun-based repositories and applications.

# Project structure
```text
<bun-repo-root>/
  AGENTS.md
  package.json
  bun.lock|bun.lockb
  biome.json
```

# Strict rules
- MUST reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST use Bun as the only runtime and package manager.
- MUST use `bun install` for dependency management.
- MUST use `bun run` for project scripts.
- MUST route test execution through Taskfile (`task validate` and `task test`); Taskfile tasks MAY invoke `bun test`.
- SHOULD prefer `bun install --backend=copyfile` for containerized installs to avoid hardlink issues.
- MUST orchestrate Bun commands through Taskfile tasks in local and CI workflows.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask for explicit `Accept` before allowing non-Bun runtime or package-manager exceptions.
