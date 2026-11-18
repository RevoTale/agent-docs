# Taskfile workflows

Use [Taskfile](https://github.com/go-task/task) as the project entry point for development, fixes, generation, validation, and tests.
Define commands in `Taskfile.yml` except where tools require their own command files. A root task orchestrates the whole project and can call tasks in nested Taskfiles.

## Task names

Use `gen`, `gen:check`, `gen:code-diff`, `fix`, `validate`, and `test` where the project has the corresponding operation. Obtain explicit maintainer approval before deviating from these interface names.

Run tests through the Taskfile. Prefer `task validate` before `task test` when both exist.

## Generated files

`task gen:check` must leave files unchanged and exit nonzero when outputs are stale.
Provide a VCS diff check for CI. When a generator lacks a dry-run mode, run `task gen:code-diff` in a clean CI checkout to regenerate outputs and check `git diff --exit-code`.

## Compose tasks

Define each reusable operation once and invoke it with `task:`. For example, both a developer and `task fix` can run the same formatter:

```yaml
version: "3"
tasks:
  format:
    cmds:
      - golangci-lint fmt ./...
  fix:
    cmds:
      - task: format
```
