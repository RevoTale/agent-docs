# Overview
This module defines baseline Taskfile workflow rules for repositories and applications.

# Project structure
```text
<repo-root>/
  AGENTS.md
  Taskfile.yml
```

# Strict rules
- Use [Taskfile](https://github.com/go-task/task) as the primary workflow entrypoint for development-related commands, including code generation, fixes, validation, and testing.
- Running Taskfile commands at repository root applies to all submodules; running inside a submodule applies only to local scope: for example, nested `task validate` should be included in the root `task validate`.
- Prefer to define workflow commands in `Taskfile.yml` instead of technology-specific files like `package.json` or `composer.json`. Keep commands in technology-specific files only when required by runtime or other app-level tools.
- Execute tests via Taskfile tasks instead of calling stack-specific test commands directly.
- Use Taskfile as the default workflows runner.
- Enforce this convention for all nested Taskfiles:
  - `task gen`: code generation committed to VCS.
  - `task gen:check`: dry-run-only generation check. Must not modify files and must return non-zero when generation would change outputs.
  - `task gen:code-diff`: runs generation and fails when a VCS/code diff check detects changes (for Git repositories: `git diff --exit-code`). Use as CI fallback when a generator has no dry-run mode.
  - `task fix`: auto-fixes issues using defined code quality tools.
  - `task validate`: runs parallel-safe tests and linters that do not mutate shared outputs.
  - `task test`: runs checks that do not fit `task validate` (for example, tests with shared mutable state).

- Keep Taskfile task naming consistent (`gen`, `gen:check`, `gen:code-diff` when needed, `fix`, `validate`, `test`). The same structure goes for the nested Taskfiles.
- If there are no tools to run for a task category, skip that category.
- When both tasks exist, run `task validate` first, then `task test`.
- Keep `task gen:check` non-mutating: only dry-run checks are allowed there.
- Every project must provide a VCS/code-diff generation check command for CI use when codegen tools do not support dry-run (for Git repositories: `git diff --exit-code`).
- Compose Taskfiles with reusable tasks and call them via `task:`. For example:
This is an incorrect variant:
```Taskfile.yml
version: "3"
tasks:
  fix:
    desc: Auto-fix lint issues
    cmds:
      - bunx biome check --write --unsafe .
```
This is a correct variant:
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
- Taskfile changes must preserve the standard task interface (`gen`, `gen:check`, `gen:code-diff` when needed, `fix`, `validate`, `test`) unless the user explicitly approves a deviation.
- Root and nested Taskfiles must remain composable so root `task` commands can orchestrate module-level workflows.
