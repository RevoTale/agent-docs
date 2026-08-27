# Overview
This module defines baseline engineering rules for TypeScript repositories and applications.

# Project structure
```text
<typescript-repo-root>/
  AGENTS.md
  package.json
  tsconfig.json|tsconfig.*.json
  biome.json|biome.jsonc
  pnpm-lock.yaml
```

# Strict rules
- MUST read and enforce matching choices from [../../awesome/javascript.md](../../awesome/javascript.md).
- MUST apply the Node.js and pnpm defaults from [../nodejs/doc.md](../nodejs/doc.md) unless a more specific runtime module applies.
- MUST reuse shared baseline rules: [../../shared/js-biome-core.md](../../shared/js-biome-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST keep TypeScript compiler options in `tsconfig.json` or `tsconfig.*.json`.
- SHOULD use a single, universal `tsconfig.json` unless tooling requires multiple configs.
- MUST configure strict compiler options, including `noUncheckedIndexedAccess`.
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST use pnpm for scripts and tooling by default; a more specific runtime module MAY override the package-manager command.
- MUST require `task validate` to include TypeScript typechecking (`pnpm exec tsc --noEmit` or the selected runtime's equivalent) and pass before merge.
- MUST generate declaration files for npm libraries that expose TypeScript types.
- SHOULD keep TypeScript configuration centralized unless tooling constraints require multiple configs.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask whether tooling constraints justify splitting `tsconfig` before approving multi-config changes.
