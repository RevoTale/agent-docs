# Overview
This module defines baseline engineering rules for React repositories and applications.

# Project structure
```text
<react-repo-root>/
  AGENTS.md
  package.json
  src/
  biome.json|biome.jsonc
  pnpm-lock.yaml
```

# Strict rules
- MUST read enforced utility/library choices from [../../awesome/react.md](../../awesome/react.md) before introducing, replacing, or removing React libraries.
- MUST enforce `required` entries from [../../awesome/react.md](../../awesome/react.md) for matching capabilities.
- MUST apply the Node.js and pnpm defaults from [../nodejs/doc.md](../nodejs/doc.md) unless a more specific runtime module applies.
- MUST reuse shared baseline rules: [../../shared/js-biome-core.md](../../shared/js-biome-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST run React build, test, and development workflows through Taskfile and pnpm by default; a more specific runtime module MAY override the package-manager command.
- SHOULD define components using functional component patterns.
- SHOULD prefer one file per component.
- MUST keep the selected package manager's lockfile in source control; use `pnpm-lock.yaml` for the default pnpm stack.
- MUST ensure React changes pass `task validate` before merge.
- MUST ensure React changes pass `task test` when the task exists.
- SHOULD colocate component-specific helpers with their component when helpers are not shared.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST present options and require `Accept` when changing component architecture conventions.
