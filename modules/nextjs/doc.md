# Overview
This module defines baseline engineering rules for Next.js repositories and applications.

# Project structure
```text
<nextjs-repo-root>/
  AGENTS.md
  app/
  next.config.ts
  biome.json|biome.jsonc
  pnpm-lock.yaml
```

# Strict rules
- MUST read enforced utility/library choices from [../../awesome/nextjs.md](../../awesome/nextjs.md) before introducing, replacing, or removing Next.js libraries.
- MUST enforce `required` entries from [../../awesome/nextjs.md](../../awesome/nextjs.md) for matching capabilities.
- MUST apply React-specific rules from [../react/doc.md](../react/doc.md).
- MUST apply the Node.js and pnpm defaults from [../nodejs/doc.md](../nodejs/doc.md) unless a more specific runtime module applies.
- MUST reuse shared baseline rules: [../../shared/js-biome-core.md](../../shared/js-biome-core.md).
- MUST reuse shared conventions: [../../shared/biome-conventions.md](../../shared/biome-conventions.md).
- MUST keep Next.js scripts executable through pnpm by default (`pnpm exec next dev`, `pnpm exec next build`, `pnpm exec next start`); a more specific runtime module MAY provide equivalent commands.
- SHOULD allow Taskfile workflows to wrap Next.js scripts for repository-level orchestration.
- MUST keep `task validate` limited to Biome checks for Next.js repositories.
- MUST keep `pnpm-lock.yaml` in source control.
- MUST avoid plain `.js` files and use `.ts` / `.tsx` instead, including configs when possible.
- MUST use App Router (`app/`) and MUST NOT introduce `pages/`.
- MUST ensure Next.js changes pass `task validate` before merge.
- MUST ensure Next.js changes pass `task test` when the task exists.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST require `Accept` for any exception to App Router-first policy.
