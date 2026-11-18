# React

Use the [TypeScript rules](../typescript/doc.md) for source files and configuration. Keep application code in `src/` unless the selected framework defines its own layout.

## Component ownership

Prefer functional components and one file per component. Keep helpers used by a single component alongside it. When changing the project's component architecture conventions, present the alternatives and obtain explicit maintainer approval.

## Tooling

Follow the [React library choices](../../awesome/react.md) for matching capabilities. Run development, build, and test workflows through Taskfile using the project's selected package manager. Commit that package manager's lockfile.
