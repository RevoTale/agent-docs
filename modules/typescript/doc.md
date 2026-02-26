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
- MUST reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST keep TypeScript compiler options in `tsconfig.json` or `tsconfig.*.json`.
- SHOULD use a single, universal `tsconfig.json` unless tooling requires multiple configs.
- MUST configure strict compiler options, including `noUncheckedIndexedAccess`.
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST run scripts and tooling through Bun via Taskfile tasks.
- MUST require `task validate` to include TypeScript typechecking (`bun run tsc --noEmit` or equivalent) and pass before merge.
- SHOULD keep TypeScript configuration centralized unless tooling constraints require multiple configs.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask whether tooling constraints justify splitting `tsconfig` before approving multi-config changes.
