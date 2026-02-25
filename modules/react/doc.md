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
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Run React build, test, and dev workflows through Bun.
- Define components using functional component patterns.
- Prefer one file per component.
- Keep Bun lockfiles in source control.

# Working Agreements
- React changes must pass repository validation or test tasks before merge.
- Keep component files focused and colocate component-specific helpers when practical.
