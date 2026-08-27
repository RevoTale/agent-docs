# Overview
This module defines the default Node.js runtime and pnpm package-management rules.

# Project structure
```text
<nodejs-repo-root>/
  AGENTS.md
  package.json
  pnpm-lock.yaml
  pnpm-workspace.yaml
```

# Strict rules
- MUST read and enforce matching choices from [../../awesome/javascript.md](../../awesome/javascript.md).
- MUST prefer the [Current Node.js release](https://nodejs.org/en/about/previous-releases) when project dependencies, tooling, and runtime environments support it.
- MUST use a supported LTS release when Current compatibility is not established or production stability requirements favor LTS.
- MUST NOT use an end-of-life Node.js release.
- MUST use [pnpm](https://pnpm.io/) as the default package manager and pin its exact version in a `package.json` field natively managed by Renovate.
- MUST use the top-level [`packageManager`](https://docs.renovatebot.com/modules/manager/npm/) field until Renovate natively manages pnpm's preferred [`devEngines.packageManager`](https://pnpm.io/package_json#devenginespackagemanager) field.
- MUST commit `pnpm-lock.yaml` and use a frozen lockfile in CI.
- MUST run package scripts through pnpm and expose repository workflows through Taskfile.
- MUST use `tsdown` for JavaScript or TypeScript libraries published to npm.
- MUST validate publishable npm libraries with `publint`; MUST also use [`attw`](https://tsdown.dev/options/package-validation) when the package publishes TypeScript declarations.
- MUST NOT migrate a working project from another runtime or package manager unless migration is explicitly in scope.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST request explicit `Accept` before selecting a non-Node.js runtime or non-pnpm package manager for a new JavaScript-family project.
