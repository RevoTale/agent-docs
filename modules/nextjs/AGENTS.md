# Overview
This module defines baseline engineering rules for Next.js repositories and applications.

# Folder Structure
```text
<nextjs-repo-root>/
  AGENTS.md
  app/
  next.config.ts
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```
- Use the React-specific rules from [../react/AGENTS.md](../react/AGENTS.md).

# Core Behaviors & Patterns
- Reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md)

# Conventions
- Reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md)
- Run Next.js workflows through Bun (`bun run next dev`, `bun run next build`, `bun run next start`).
- Keep Bun lockfiles in source control.
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Use only `app` directory features.

# Working Agreements
- `bun run next build` must pass for Next.js changes before merge.
- Next.js changes must use the App Router (`app/`) and must not introduce `pages/`.
