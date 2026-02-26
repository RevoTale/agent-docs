# Overview
This module defines baseline Taskfile workflow rules for repositories and applications.

# Project structure
```text
<repo-root>/
  AGENTS.md
  Taskfile.yml
```

# Strict rules
- MUST use [Taskfile](https://github.com/go-task/task) as the primary workflow entrypoint for development commands, including code generation, fixes, validation, and testing.
- MUST treat Taskfile execution from repository root as project-wide orchestration; nested Taskfiles MUST be composable from root tasks.
- MUST define workflow commands in `Taskfile.yml` unless a runtime or tool requires technology-specific command files.
- MUST execute tests via Taskfile tasks instead of direct stack-specific test commands.
- MUST keep task interface names consistent as `gen`, `gen:check`, `gen:code-diff` (when needed), `fix`, `validate`, and `test`.
- MUST keep `task gen:check` non-mutating and return non-zero when generation would change outputs.
- MUST provide `task gen:code-diff` as CI fallback when generators do not support dry-run checks.
- SHOULD run `task validate` before `task test` when both tasks exist.
- MAY skip a task category when no relevant tools exist.
- MUST provide a VCS/code-diff generation check for CI use (for Git repositories: `git diff --exit-code`).
- MUST compose reusable tasks and invoke them via `task:`.

Example of incorrect composition:
```Taskfile.yml
version: "3"
tasks:
  fix:
    desc: Auto-fix lint issues
    cmds:
      - bunx biome check --write --unsafe .
```

Example of correct composition:
```Taskfile.yml
version: "3"
tasks:
  biome:fix:
    desc: Fix files with Biome
    cmds:
      - bunx biome check --write --unsafe .
  fix:
    desc: Auto-fix all code quality issues
    cmds:
      - task: biome:fix
```

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask for explicit `Accept` before approving deviations from standard task interface names.
