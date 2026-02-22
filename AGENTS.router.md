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
    go/AGENTS.md       # load for Go repositories/services
    nextjs/AGENTS.md   # load for Next.js repositories/apps
```

# Working Agreements
- Always load `modules/common/AGENTS.md`.
- Load `modules/go/AGENTS.md` only when at least one signal exists: `go.mod`, `go.work`, `*.go`, `cmd/`, `internal/`.
- Load `modules/nextjs/AGENTS.md` only when at least one signal exists: `next.config.js|mjs|ts`, `package.json` with `next`, `app/` containing route files, or `pages/` with Next.js conventions.
- If signals for multiple stacks exist, load all matching stack modules.
- If no stack signal matches, keep only `modules/common/AGENTS.md` plus any local repository `AGENTS.md`.
- Precedence order for conflicts: nearest local `AGENTS.md` in target repository, then repository root `AGENTS.md`, then `AGENTS.router.md`, then loaded modules.
- Any module change must be specific, testable, and paired with an update in `AGENTS.router.md` when load conditions change.
