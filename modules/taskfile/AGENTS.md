# Overview
This module defines baseline Taskfile workflow rules for repositories and applications.

# Folder Structure
```text
<repo-root>/
  AGENTS.md
  Taskfile.yml
```

# Core Behaviors & Patterns
- Use Taskfile as the primary workflow entrypoint for code generation, fixes, and validation.
- Running Taskfile commands at repository root applies to all submodules; running inside a submodule applies only to local scope.
- Prefer to have all commands written in the `Taskfile.yml` instead of in technology specific composing files like `package.json`, `composer.json` etc. Place command in technology specific files only when they are required for the technologuyn features.

# Conventions
- Use Taskfile for code generation, fixes, and validation.
- Enforce the strict convetion for all nested Taskfiles:
  - `task gen` - Code generation. Required to run before testing/merging to keep. Should container code generators that are being run during development and commited to the VCS.
  - `task gen:check` - Should check whether  code generation was not forgotten to run. It should implement via built-in generator script "dry run", or via "git diff" or similar commands and return an 1 error code, if code actually changes.
  - `task fix` - should auto fix with the all predefined code quality tool.
  - `task validate` - should run all test/linters/code quality tools, that do not require specific runtime work. `task validate` should only include parallel-safe checks that do not touch shared outputs. If a check is not parallel-friendly or mutates shared files, it belongs under `test` instead. Configure them to run in parallel.
  - `task test` - should run all the remaining tests that did not fit the `task validate` requirements. Commonly they are unit/integration tests.

# Working Agreements
- Keep Taskfile task naming consistent (`gen`, `gen:check`, `fix`, `validate`, `test`). The same structure goes for the nested Taskfiles.
- If some script has no utilities to run, skip them. 
- Compose taskfile to have the reusable separated command, then compose it via `task: ` property. For example:
This is INCORRECT variant:
```Taskfile.yml
version: "3"
  fix:
    desc: Auto-fix lint issues
    cmds:
      - bunx biome check --write --unsafe .
```
This is a Corrent variant:
```Taskfile.yml
version: "3"
  biome:fix:
    desc: Fix the files with Biome
    cmds:
      - bunx biome check --write --unsafe .
  fix:
    desc: Auto-fix all code quality issues
    cmds:
      - task: biome:fix
```
