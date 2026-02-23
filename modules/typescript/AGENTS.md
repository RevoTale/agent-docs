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
- Prefer a single, universal `tsconfig.json` when practical.
- Configure strict compiler options, including `noUncheckedIndexedAccess`.
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Run scripts and tooling through Bun.

# Working Agreements
- `bun run tsc --noEmit` (or equivalent typecheck task) must pass before merge.
- Keep TypeScript configuration centralized unless multiple configs are required by tooling constraints.
