# Overview
This file is the universal AGENTS router and composition contract.
It defines which module files to load and how to merge them into a target repository `AGENTS.md`.

# Canonical Module Registry
Use this table as the single source of truth for module routing.

| key | full_name | module_path | load_when |
| --- | --- | --- | --- |
| common | Common baseline rules | modules/common/AGENTS.md | always |
| taskfile | Taskfile (go-task) workflows | modules/taskfile/AGENTS.md | always |
| typescript | TypeScript | modules/typescript/AGENTS.md | `tsconfig.json`, `tsconfig.*.json`, `*.ts`, or `*.tsx` exist |
| react | React | modules/react/AGENTS.md | `package.json` includes `react`, or `*.jsx` / `*.tsx` files exist |
| nextjs | Next.js | modules/nextjs/AGENTS.md | `next.config.js|mjs|ts`, `package.json` includes `next`, or `app/` route files exist |
| bun | Bun runtime/package manager | modules/bun/AGENTS.md | `bun.lock` / `bun.lockb` / `bunfig.toml` exist, or `package.json` uses Bun tooling |
| go | Go | modules/go/AGENTS.md | `go.mod`, `go.work`, `*.go`, `cmd/`, or `internal/` exist |

# Routing Rules
1. Load all rows where `load_when` is `always`.
2. Evaluate all other rows and load a module when any signal in its `load_when` condition matches.
3. If multiple modules match, load all matched modules.
4. Any module add/remove/rename or signal change must update this table in the same change.

# Instructions for combining the rules
Each module must define the following sections:
- Overview
- Project structure
- Strict rules
- Working Agreements

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
1. Default behavior is additive merge.
2. Stack modules have highest priority.
3. Compatible requirements must be merged together.
Example: if one rule requires `task validate` to include `tsc --noEmit` and another requires `golangci-lint`, include both commands.
4. If strict rules still conflict (`must` vs `must not`) and cannot be merged, the more specific stack wins.
Example: `nextjs` overrides `react` on overlapping frontend rules.
5. If conflicting stacks are equally specific, switch to interview mode for user resolution.
6. Handle these cases in interview mode with the user:
- identical duplicated rules that could be deduplicated
- truly incompatible rules (`must` vs `must not` on the same behavior)

### Working Agreements
This part must be agreed with the user.
First merge Working Agreements from loaded modules and propose the result.
The user must give a clear response of `Accept` or request edits.
Iterate until consensus and finalize only after explicit `Accept`.
