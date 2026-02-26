# Overview
This module defines baseline engineering rules for Next.js repositories and applications.

# Project structure
```text
<nextjs-repo-root>/
  AGENTS.md
  app/
  next.config.ts
  biome.json|biome.jsonc
  bun.lock|bun.lockb
```

# Strict rules
- MUST apply React-specific rules from [../react/doc.md](../react/doc.md).
- MUST reuse shared baseline rules: [../../shared/js-biome-bun-core.md](../../shared/js-biome-bun-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST keep Next.js scripts executable through Bun (`bun run next dev`, `bun run next build`, `bun run next start`) because Dockerfile builds depend on them.
- SHOULD allow Taskfile workflows to wrap Next.js scripts for repository-level orchestration.
- MUST keep `task validate` limited to Biome checks for Next.js repositories.
- MUST keep Bun lockfiles in source control.
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST use App Router (`app/`) and MUST NOT introduce `pages/`.
- MUST ensure Next.js changes pass `task validate` before merge.
- MUST ensure Next.js changes pass `task test` when the task exists.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST require `Accept` for any exception to App Router-first policy.
