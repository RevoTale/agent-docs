# Overview
This is the universal AGENTS router for the organization. It selects stack modules by repository signals so agent context is loaded lazily.

# Folder Structure
Module index:

```text
<repo-root>/
  AGENTS.md            # stable bootstrap entrypoint
  AGENTS.router.md     # routing and load conditions
  modules/
    common/AGENTS.md   # always loaded
    taskfile/AGENTS.md # Taskfile (go-task) workflow module
    typescript/AGENTS.md # TypeScript repositories/apps
    react/AGENTS.md     # React repositories/apps
    nextjs/AGENTS.md    # Next.js repositories/apps
    bun/AGENTS.md       # Bun runtime/package manager repositories/apps
    go/AGENTS.md        # Go repositories/services
```

# Working Agreements
- Always load `modules/common/AGENTS.md`.
- Always load `modules/taskfile/AGENTS.md` (short key: `taskfile`, full stack name: `Taskfile (go-task)`).
- Load `modules/typescript/AGENTS.md` when at least one signal exists: `tsconfig.json`, `tsconfig.*.json`, `*.ts`, or `*.tsx`.
- Load `modules/react/AGENTS.md` when at least one signal exists: `package.json` with `react`, `*.jsx`, or `*.tsx`.
- Load `modules/go/AGENTS.md` only when at least one signal exists: `go.mod`, `go.work`, `*.go`, `cmd/`, `internal/`.
- Load `modules/nextjs/AGENTS.md` only when at least one signal exists: `next.config.js|mjs|ts`, `package.json` with `next`, `app/` containing route files, or `pages/` with Next.js conventions.
- Load `modules/bun/AGENTS.md` when at least one Bun signal exists: `bun.lock`, `bun.lockb`, `bunfig.toml`, or `package.json` with Bun scripts/tooling.
- If JavaScript, TypeScript, React, or Next.js signals exist, also load `modules/bun/AGENTS.md`.
- If signals for multiple stacks exist, load all matching stack modules.
- If no stack signal matches, keep `modules/common/AGENTS.md` and `modules/taskfile/AGENTS.md` plus any local repository `AGENTS.md`.
- Precedence order for conflicts: nearest local `AGENTS.md` in target repository, then repository root `AGENTS.md`, then `AGENTS.router.md`, then loaded modules.
- Any module change must be specific, testable, and paired with an update in `AGENTS.router.md` when load conditions change.
