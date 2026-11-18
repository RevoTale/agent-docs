# Standards for every project

## Development workflow

Use the [Taskfile workflow](../taskfile/doc.md) as the project entry point.
Before merge, `task validate` must pass, along with `task test` when the project defines it.

## Dependencies and environments

Keep `renovate.json` at the repository root and follow the [dependency policy](../../shared/dependency-management.md) when adding or updating packages and images.
Use the [OS and base image policy](../../shared/os-and-base-images.md) when choosing or changing an environment.

## Library choices

Consult the project's selected library tables before introducing, replacing, or removing a library.
Each row applies only when the project needs that capability. Use the listed package for a `required` row,
prefer the listed package for a `preferred` row, and exclude a `banned` package.
For a deviation from a required choice, obtain explicit maintainer approval and record its scope and reason in the project.
