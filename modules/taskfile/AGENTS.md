# Overview
This module defines baseline Taskfile workflow rules for repositories and applications.

# Folder Structure
```text
<repo-root>/
  AGENTS.md
  Taskfile.yml
```

# Core Behaviors & Patterns
- Use Taskfile as the primary workflow entrypoint for code generation, fixes, validation, and testing.
- Running Taskfile commands at repository root applies to all submodules; running inside a submodule applies only to local scope.
- Prefer to define workflow commands in `Taskfile.yml` instead of technology-specific files like `package.json` or `composer.json`. Keep commands in technology-specific files only when required by those tools.
- Execute tests via Taskfile tasks instead of calling stack-specific test commands directly.

# Conventions
- Use Taskfile for code generation, fixes, validation, and testing.
- Enforce this convention for all nested Taskfiles:
  - `task gen`: code generation committed to VCS.
  - `task gen:check`: verifies generation is up to date and returns non-zero when generated outputs differ.
  - `task fix`: auto-fixes issues using defined code quality tools.
  - `task validate`: runs parallel-safe tests and linters that do not mutate shared outputs.
  - `task test`: runs checks that do not fit `task validate` (for example, tests with shared mutable state).

# Working Agreements
- Keep Taskfile task naming consistent (`gen`, `gen:check`, `fix`, `validate`, `test`). The same structure goes for the nested Taskfiles.
- If there are no tools to run for a task category, skip that category.
- When both tasks exist, run `task validate` first, then `task test`.
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
