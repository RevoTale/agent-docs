---
name: init-project-from-agent-docs
description: Initialize or refresh a repository AGENTS.md by fetching the latest router and module policies from https://github.com/RevoTale/agent-docs at runtime, then linking only the modules that match the target project. Use for new repo bootstrap, AGENTS migration, or policy sync when hardcoded rules must be avoided.
---

# Init Project From Agent Docs

Use this workflow for Codex-compatible skill execution.

## Workflow

1. Fetch the latest source from the web first.
- Get the current default branch for `RevoTale/agent-docs` from GitHub API.
- Fetch `AGENTS.router.md` from that branch.
- Discover module paths from router text (do not assume fixed module names).
- Fetch module files only as needed for stack matching and validation.
- Never rely on memory or stale local copies of this repository.

2. Detect target repository signals.
- Inspect the target repository files and folders.
- Evaluate module load conditions from the fetched router text.
- Select modules strictly from router conditions.
- If router includes an always-load module rule, include it.

3. Create or update target root `AGENTS.md`.
- Keep root content repository-specific.
- Add a `Base Policy Links (Load First)` section.
- Add links to:
  - `AGENTS.router.md`
  - any selected module files
- Place local repository rules in a separate `Local Details` section below the base links.

4. Refresh existing repositories safely.
- Preserve existing local details unless they conflict with current repo reality.
- Replace outdated base links with links to the latest `RevoTale/agent-docs` branch (or pin to a commit when requested).
- Do not copy policy text inline when links are sufficient.

5. Validate before finishing.
- Confirm each linked file exists on GitHub.
- Confirm module selection is justified by router conditions.
- Confirm no hardcoded stack rules were introduced by the skill itself.
- Report exactly which remote files were used.

## Compatibility Notes

- Codex: this `SKILL.md` is the primary entrypoint.
- Claude: use the mirrored instructions in `CLAUDE.md` in the same directory.
