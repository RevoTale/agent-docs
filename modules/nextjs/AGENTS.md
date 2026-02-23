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
- Keep Next.js scripts executable through Bun (`bun run next dev`, `bun run next build`, `bun run next start`) because Dockerfile builds depend on them.
- Taskfile workflows may wrap these scripts for repository-level orchestration.
- For Next.js repositories, keep `task validate` limited to Biome checks.
- Keep Bun lockfiles in source control.
- Avoid plain `.js` files. Only `.tsx` and `.ts`, even for configs, if possible.
- Use only `app` directory features.

# Working Agreements
- `task validate` and `task test` must pass for Next.js changes before merge.
- Next.js changes must use the App Router (`app/`) and must not introduce `pages/`.
