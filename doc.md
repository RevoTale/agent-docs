# Overview
This is the universal AGENTS router. It selects stack modules by repository signals so agent can find them and combine to compose the own `AGENTS.md` speific to the use project.

# Folder Structure
Module index:

```text
<repo-root>/
  AGENTS.md            # stable bootstrap entrypoint
  doc.md            # instructions on composing the target prpject AGENTS.md
  modules/
    common/doc.md   # always loaded
    taskfile/doc.md # Taskfile (go-task) workflow module
    typescript/doc.md # TypeScript repositories/apps
    react/doc.md     # React repositories/apps
    nextjs/doc.md    # Next.js repositories/apps
    bun/doc.md       # Bun runtime/package manager repositories/apps
    go/doc.md        # Go repositories/services
```

# Examples of common stack-specific signals
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
- `modules/nextjs/doc.md`:
  - Files like `next.config.js|mjs|ts`, `package.json` with `next`, or `app/` containing route files.
- `modules/bun/doc.md`:
  - Files like `bun.lock`, `bun.lockb`, `bunfig.toml`, or `package.json` with Bun scripts/tooling.
  - JavaScript, TypeScript, React, or Next.js signals exist.

# Change management:
Any module change must be specific, testable, and paired with an update in `doc.md` when load conditions change.

# Stack matching behavior:
If signals for multiple stacks exist, load all matching stack modules.


# Instructions for combining the rules

Each module strictly defines the following sections:
- Overview
- Strict rules
- Project structure
- Working Agreements

## Rules of the sections combining into a single AGENTS.md
We have a strict rules how to combine those section. 
Following merge rules should be followed when composing the stack specific modules into single `AGENTS.md`. They are described by sections, starting with the `###`.


### Overview 
Modules description. Needed for better understanding of the module specifics by the agent.

### Project structure
Here is an example format of the project strcture section.

```text
<go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
```

- `OR` condition is marked as `|`
- Variable naming of files/folder are defined via `<variadic-description>` where `variadic-description` is the any short text that describes the naming. 
- Folder name are ending via `/` symbol and the nesting level is marked via the space count used. Tabulation ending should be marked as a `|` symbol. For example, the 2-level nested folder will marked as `  |`. No single `|` can be used beacuse root folders are marked without any spaces and symbold.

When merging two folder structures, merge them using rules above.

#### Example of mering two Project structures
##### Strcture example 1
```text
<go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
```
##### Strcture example 2
```text
<bun-repo-root>/
 |AGENTS.md
 |package.json
 |biome.json
 |bun.lock|bun.lockb
```
#####  Merged strcture examples 1+2
```text
<bun-go-repo-root>/
 |AGENTS.md
 |.golangci.yml|.golangci.yaml|.golangci.toml
 |package.json
 |biome.json
 |bun.lock|bun.lockb
```

### Strict rules
Strict rules should be used as is. Just add one to another. In case of conflicting rules inteview the user on how better to resolve the conflict.

### Working Agreements

This part must be agreed upon with the user. First, merge the Working Agreement from the existing modules and propose the final version to the user. The user must give a clear response of “Accept” or indicate what they are not satisfied with and provide their own edits. After that, propose a new version considering their suggestions. Continue until you reach consensus and they respond “Accept” to the final proposed version.