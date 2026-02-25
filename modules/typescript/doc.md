# Overview
This module defines baseline engineering rules for TypeScript repositories and applications.

# Project structure
```text
<typescript-repo-root>/
  AGENTS.md
  package.json
  tsconfig.json|tsconfig.*.json
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```

# Strict rules
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Keep TypeScript compiler options in `tsconfig.json` or `tsconfig.*.json`.
- Prefer a single, universal `tsconfig.json` when practical.
- Configure strict compiler options, including `noUncheckedIndexedAccess`.
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Run scripts and tooling through Bun via Taskfile tasks.

# Working Agreements
- `task validate` must include TypeScript typechecking (`bun run tsc --noEmit` or equivalent) and pass before merge.
- Keep TypeScript configuration centralized unless multiple configs are required by tooling constraints.
