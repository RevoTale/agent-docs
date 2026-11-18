# Node.js and pnpm

Use Node.js and pnpm for a new JavaScript-family project. Selecting a different runtime or package manager requires explicit maintainer approval. For an existing project, migrate its runtime or package manager only when the change explicitly includes that migration.

## Node.js version

Prefer the [Current Node.js release](https://nodejs.org/en/about/previous-releases) when dependencies, tooling, and the runtime environment support it. Use a supported LTS release when compatibility is uncertain or production stability favors LTS. Use a release that still receives support.

## Package management

Pin the exact [pnpm](https://pnpm.io/) version in a field Renovate manages natively. Use the top-level [`packageManager` field](https://docs.renovatebot.com/modules/manager/npm/) until Renovate natively manages pnpm's preferred [`devEngines.packageManager` field](https://pnpm.io/package_json#devenginespackagemanager). Configure dependency age in `pnpm-workspace.yaml` as required by the [dependency policy](../../shared/dependency-management.md).

Commit `pnpm-lock.yaml` and use a frozen lockfile in CI. Run package scripts through pnpm and expose repository workflows through Taskfile.
