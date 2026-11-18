# RevoTale engineering standard

This is the engineering standard for RevoTale projects. People and coding agents read the same Markdown. Each project keeps the part it uses in Git under `.revotale-agent-doc/`. Its root `AGENTS.md` links directly to those files and names the project's validation commands.

## Use it in a project

- **Start a project.** Choose the architecture from the product requirements. Select the baseline, the modules for each app or service, and the library choices for the capabilities it needs. Apply the standard from the first change.
- **Adopt it in an existing project.** Identify the current architecture, the intended architecture and the work in scope. New work follows the selected rules within its boundary. Record a gap when meeting a rule requires a wider migration. Agree on that migration before changing runtimes, tools or public behavior.
- **Make a change.** Start in the nearest applicable `AGENTS.md`. Read the local rules for the decision at hand, make the change and run the project's checks. Reopen the affected rule when the task changes scope, the result conflicts with a rule or context is lost.

The [adoption guide](examples/adoption/README.md) walks through a Go service, including its vendir configuration, project bootstrap and upgrade from the earlier remote links. The optional [apply-agent-docs skill](skills/apply-agent-docs/SKILL.md) guides selection and migration.

## Choose the modules

Select modules per app, service or package. Manifests and lockfiles identify the current stack. For a new boundary, use the agreed target. Keep repository-wide guidance in the root bootstrap and add a nested bootstrap where the rules differ.

| Module | When it applies |
| --- | --- |
| [common](modules/common/doc.md) | Every project. Includes OS, base image and dependency policies. |
| [taskfile](modules/taskfile/doc.md) | Every project. Defines the shared task interface. |
| [nodejs](modules/nodejs/doc.md) | New JavaScript-family projects by default, and boundaries using Node.js with pnpm. |
| [bun](modules/bun/doc.md) | Boundaries using Bun, or new ones with an approved Bun exception. |
| [typescript](modules/typescript/doc.md) | TypeScript source or compiler configuration. Select the runtime separately. |
| [react](modules/react/doc.md) | React components or a React dependency. Include TypeScript. |
| [nextjs](modules/nextjs/doc.md) | A Next.js app. Include React, TypeScript and the applicable runtime. |
| [go](modules/go/doc.md) | A Go service, module or workspace. |
| [templ](modules/templ/doc.md) | Go code rendering HTML with templ. Include Go and the paired templ examples. |

Use the [library registry](awesome/index.md) for choices that depend on a capability, such as authentication or publishing to npm. A required entry applies when the project uses that capability. Include the selected registry files and local documents linked by the selected modules. The central catalog and contributor instructions stay here.

## Apply a rule

A rule applies within its stated scope. A narrower rule resolves a conflict with a general default. Record an approved exception in the affected project's bootstrap with its scope and reason. When two applicable requirements still conflict, bring that decision to the maintainer before making the affected change.

Read paired examples when distinguishing two implementations is difficult. Put mechanically enforceable requirements in the project's checks where practical. Review covers architectural choices that require judgment.

## Update a project's copy

Choose a reviewed revision, sync the selected files, and review the diff together with the lockfile and bootstrap. Ordinary work reads the committed copy offline. The revision identifies the source bytes. Review establishes whether changes to those instructions are appropriate for the project.

For changes to the standard itself, read the [maintainer guide](AGENTS.md).
