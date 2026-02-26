# Overview
This file is the universal AGENTS router and composition contract.
It defines which module files to load and how to merge them into a target repository `AGENTS.md`.

# Canonical Module Registry
Use this table as the single source of truth for module routing.

| key | full_name | module_path | load_when |
| --- | --- | --- | --- |
| common | Common baseline rules | modules/common/doc.md | always |
| taskfile | Taskfile (go-task) workflows | modules/taskfile/doc.md | always |
| typescript | TypeScript | modules/typescript/doc.md | `tsconfig.json`, `tsconfig.*.json`, `*.ts`, or `*.tsx` exist |
| react | React | modules/react/doc.md | `package.json` includes `react`, or `*.jsx` / `*.tsx` files exist |
| nextjs | Next.js | modules/nextjs/doc.md | `next.config.js|mjs|ts`, `package.json` includes `next`, or `app/` route files exist |
| bun | Bun runtime/package manager | modules/bun/doc.md | `bun.lock` / `bun.lockb` / `bunfig.toml` exist, or `package.json` uses Bun tooling |
| go | Go | modules/go/doc.md | `go.mod`, `go.work`, `*.go`, `cmd/`, or `internal/` exist |

# Routing Rules
1. Load all rows where `load_when` is `always`.
2. Evaluate all other rows and load a module when any signal in its `load_when` condition matches.
3. If multiple modules match, load all matched modules.
4. Any module add/remove/rename or signal change must update this table in the same change.

# Instructions for combining the rules
Each module must define the following sections in this exact order:
- Overview
- Project structure
- Strict rules
- Working Agreements

Section semantics:
- `Strict rules` MUST contain technical requirements, commands, checks, and invariants.
- `Working Agreements` MUST describe user-agent interaction protocol only and MUST NOT include technical command/check constraints.

## Rules of sections combining into a single AGENTS.md
Follow these merge rules by section.

### Overview
Keep module descriptions for context and stack intent. Merge without removing relevant stack context.

### Project structure
Use this format:

```text
<go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
```

- `OR` condition is marked as `|`.
- Variable naming is defined via `<variadic-description>`.
- Folder names end with `/`.
- Keep indentation stable and represent structure consistently.

When merging two project structures, produce a union of paths and preserve `OR` groups.

#### Example of merging two Project structures
##### Structure example 1
```text
<go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
```
##### Structure example 2
```text
<bun-repo-root>/
 |AGENTS.md
 |package.json
 |biome.json
 |bun.lock|bun.lockb
```
##### Merged structure examples 1+2
```text
<bun-go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
 |package.json
 |biome.json
 |bun.lock|bun.lockb
```

### Strict rules
- MUST merge compatible requirements additively.
- MUST treat stack modules as higher priority than baseline modules when strict rules conflict.
- MUST let the more specific stack win when strict rules are incompatible and cannot be merged.
- MUST switch to interview mode when conflicting stacks are equally specific.
- MUST use interview mode to decide whether identical duplicated rules should be deduplicated.
- MUST use interview mode for truly incompatible rules (`must` vs `must not` on the same behavior).
- MUST document user-approved technical exceptions directly in `Strict rules`.

### Working Agreements
- MUST propose merged `Working Agreements` to the user before finalizing.
- MUST wait for explicit `Accept` before finalizing policy text.
- MUST iterate until consensus and finalize only after explicit `Accept`.

# Enforcement
## Checklist
- MUST update the canonical module registry whenever module paths or load signals change.
- MUST keep module section order and names exactly as specified.
- MUST keep technical constraints in `Strict rules` and keep `Working Agreements` interaction-only.
- MUST run `go run ./scripts/validate-agent-docs.go` before merge.

## CI Validation
- MUST keep `.github/workflows/validate-agent-docs.yml` active so policy validation runs on pushes and pull requests.
