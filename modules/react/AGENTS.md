# Overview
This module defines baseline engineering rules for React repositories and applications.

# Folder Structure
```text
<react-repo-root>/
  AGENTS.md
  package.json
  src/
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```

# Core Behaviors & Patterns
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)

# Conventions
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Run React build, test, and dev workflows through Bun.
- For the components definition use only functional components format. 
- Prefer one file per component.
- Keep Bun lockfiles in source control.

# Working Agreements
