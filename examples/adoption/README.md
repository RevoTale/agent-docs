# Adopt the standard in a project

This example covers a Go service. Contributors can read the checked-in Markdown with a text editor. Updating the copy uses Git and [vendir](https://carvel.dev/vendir/). The snapshot check runs with Go in a trusted checkout of this standard, including for projects built with another language.

## Select the documents

Start with the [catalog](../../README.md). Select the baseline, the modules for each app or service, and the library choices for the capabilities it actually uses. Include the relative links those documents depend on. Keep the union in one snapshot and route each subtree to its own modules.

| Project decision | Selection beyond the baseline |
| --- | --- |
| Go service | Go module and Go library choices, as in the example config. |
| Go service adds templ views | Add the templ module and its linked Markdown examples under `examples/templ/`. |
| New TypeScript/React app | Select Node.js/pnpm, TypeScript and React. Include their linked Biome guidance and library choices. |
| Existing TypeScript/React app on Bun | Select Bun as the runtime alongside TypeScript and React. Add Next.js only for a Next.js app. |
| Publish an npm package or add authentication | Include the applicable `awesome/javascript.md` or `awesome/auth.md` entry in the selection and link it from the owning scope. |

## Copy the reviewed revision

1. Copy [vendir.yml](vendir.yml) to the project root. If the project already uses vendir, add the example's directory entry to its existing configuration. Reserve `.revotale-agent-doc/` for synchronized files. Move any project-authored notes elsewhere before the first sync.
2. Replace `REPLACE_WITH_REVIEWED_COMMIT_SHA` with the full SHA of a reviewed upstream commit containing the snapshot-based standard. List the chosen files in `includePaths` and retain the upstream license. Review the source URL and paths before syncing.
3. From the project root, run:

   ```sh
   vendir sync --directory .revotale-agent-doc
   ```

   vendir writes the selected files under `.revotale-agent-doc/` and records the resolved SHA in `vendir.lock.yml`. The source config and lockfile live outside the managed directory.

## Point contributors to the local copy

Merge a short section like this into the service's root `AGENTS.md`:

```markdown
## Engineering standard

Read the [common rules](.revotale-agent-doc/modules/common/doc.md) and
[Taskfile workflows](.revotale-agent-doc/modules/taskfile/doc.md).
For Go changes, use the [Go module](.revotale-agent-doc/modules/go/doc.md).
When choosing a Go library, consult the [Go library list](.revotale-agent-doc/awesome/go.md).
Run `task validate` and `task test` for this service.
```

For an existing project, replace the old `blob/main` policy links in root and nested `AGENTS.md` files. Preserve project commands, ownership boundaries, accepted exceptions and user-authored instructions. Adjust link paths relative to each file. Describe any gap between the current stack and the target in the project’s existing architecture notes. Schedule migrations within their agreed scope.

## Review an update

Change the pinned SHA or the selection, then run the same sync command. For a Go service, inspect the Go rules, library choices and both baseline modules before adopting the update.

1. From a trusted checkout of this standard, check the consumer's local links and file types:

   ```sh
   go run ./scripts/validate-agent-docs.go -snapshot /absolute/path/to/project/.revotale-agent-doc
   ```

2. Open the relative links from each changed root or nested `AGENTS.md` in the consumer. Stage the config, lockfile, snapshot and affected bootstrap files together. Review the staged diff, including new files and their modes. Check that the locked SHA matches the reviewed revision and that every copied file belongs to the selection.
3. Reproduce the staged snapshot from the lock and check for drift in the consumer:

   ```sh
   vendir sync --locked --directory .revotale-agent-doc
   git diff --exit-code -- vendir.lock.yml .revotale-agent-doc
   git ls-files --others -- .revotale-agent-doc
   ```

   The last command should print nothing. Run the project's checks for any changes to code or tooling. Ordinary development reads the committed files offline. The pin identifies the approved revision, while review determines whether its instructions belong in the project.
