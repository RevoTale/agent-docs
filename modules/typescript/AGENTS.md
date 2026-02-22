# Overview
This module defines baseline engineering rules for TypeScript repositories and applications.

# Folder Structure
```text
<typescript-repo-root>/
  AGENTS.md
  package.json
  tsconfig.json|tsconfig.*.json
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```

# Core Behaviors & Patterns
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)

# Conventions
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Keep TypeScript compiler options in `tsconfig.json` or `tsconfig.*.json`.
- Prefer the only one, unversal `tsconfig.json` as a config.
- `tsconfig.json` shpudl be configured to have the strictest rules: including the unsafe array ketys access.
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.

- Run scripts and tooling through Bun.

# Working Agreements
