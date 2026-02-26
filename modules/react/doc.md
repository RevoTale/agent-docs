# Overview
This module defines baseline engineering rules for React repositories and applications.

# Project structure
```text
<react-repo-root>/
  AGENTS.md
  package.json
  src/
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```

# Strict rules
- MUST reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST run React build, test, and development workflows through Bun.
- SHOULD define components using functional component patterns.
- SHOULD prefer one file per component.
- MUST keep Bun lockfiles in source control.
- MUST ensure React changes pass `task validate` before merge.
- MUST ensure React changes pass `task test` when the task exists.
- SHOULD colocate component-specific helpers with their component when helpers are not shared.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST present options and require `Accept` when changing component architecture conventions.
