# Bun

Apply this module where a working project already requires Bun or maintainers explicitly select it. Obtain explicit approval before migrating that project to another runtime or package manager.

## Packages and workflows

Use Bun as the only runtime and package manager within this project boundary. Use `bun install` for dependencies and `bun run` for scripts. In containers, prefer `bun install --backend=copyfile` when hardlinks cause installation problems.

Expose local and CI workflows through Taskfile. Use `task validate` and `task test` as project entry points. The underlying test task may invoke `bun test`.

## Code quality

Apply the [Biome baseline](../../shared/js-biome-core.md) and [configuration conventions](../../shared/biome-conventions.md).
