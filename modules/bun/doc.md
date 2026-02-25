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
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Use Bun as the only runtime and package manager.
- Use `bun install` for dependency management.
- Use `bun run` for project scripts.
- Route test execution through Taskfile (`task validate` and `task test`); Taskfile tasks may invoke `bun test`.
- For containerized installs, prefer `bun install --backend=copyfile` to avoid hardlink issues on bind mounts and Dev Containers host/bin volume synchronization for `node_modules`.

# Working Agreements
- Bun commands must be orchestrated through Taskfile tasks in local and CI workflows.
