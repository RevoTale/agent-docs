# Overview
This module defines baseline engineering rules for Bun-based repositories and applications.

# Folder Structure
```text
<bun-repo-root>/
  AGENTS.md
  package.json
  bun.lock|bun.lockb
  biome.json
```

# Core Behaviors & Patterns
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)
- Use Bun as the only runtime and package manager.

# Conventions
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Use `bun install` for dependency management.
- Use `bun run` for project scripts.
- Use `bun test` for test execution when tests are available.

# Working Agreements
- Bun commands must be used in local and CI workflows.
