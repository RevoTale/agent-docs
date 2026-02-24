# Overview
This is the universal AGENTS router. It selects stack modules by repository signals so agent can find them and combine to compose the own `AGENTS.md` speific to the use project.

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

# Instructions
Here are the stacks defined and some common signals on when to use them.

The format is:
```
- <technology stack>
  - Common signal 1
  - Common signal 2
```

- `modules/taskfile/doc.md`:
  - use always.
- `modules/typescript/doc.md`:
  - Any files `tsconfig.json`, `tsconfig.*.json`, `*.ts`, or `*.tsx`
- `modules/react/doc.md`: 
  - Files like `package.json` with `react` or `*.jsx`, or `*.tsx`.
- `modules/go/doc.md`:
  - Any files like `go.mod`, `go.work`, `*.go`, `cmd/`, `internal/`.




### the remining stack to refactor
- Load `modules/nextjs/doc.md` (short key: `nextjs`, full stack name: `Next.js`) only when at least one signal exists: `next.config.js|mjs|ts`, `package.json` with `next`, or `app/` containing route files.
- Load `modules/bun/doc.md` (short key: `bun`, full stack name: `Bun`) when at least one Bun signal exists: `bun.lock`, `bun.lockb`, `bunfig.toml`, or `package.json` with Bun scripts/tooling.
- If JavaScript, TypeScript, React, or Next.js signals exist, also load `modules/bun/doc.md`.
- If signals for multiple stacks exist, load all matching stack modules.
- If no stack signal matches, keep `modules/common/doc.md` and `modules/taskfile/doc.md` plus any local repository `AGENTS.md`.
- Precedence order for conflicts: nearest local `AGENTS.md` in target repository, then repository root `AGENTS.md`, then `AGENTS.router.md`, then loaded modules.
- Any module change must be specific, testable, and paired with an update in `AGENTS.router.md` when load conditions change.
