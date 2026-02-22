---
name: init-project-from-agent-docs
description: Create or refresh a repository AGENTS.md by fetching the latest router and module policies from https://github.com/RevoTale/agent-docs at runtime, identifying project technology patterns, and linking only matching modules. Use when AGENTS.md is missing or outdated and hardcoded stack rules must be avoided.
---

# Init Or Refresh AGENTS.md From Agent Docs

Use this workflow with any AI agent.

## Workflow

1. Fetch the latest source of truth from the web.
- Get the current default branch for `RevoTale/agent-docs` from GitHub API.
- Fetch `AGENTS.router.md` from that branch.
- Discover available module paths from router text and/or `modules/` directory listing.
- Do not assume fixed technology names or a fixed module set.
- Fetch module files only as needed for matching and validation.
- Never rely on memory or stale local copies of this repository.

2. Learn target repository patterns.
- Inspect the target repository files and folders.
- Collect concrete signals (manifests, lockfiles, source extensions, framework folders, build config).
- Evaluate router load conditions against those signals.
- If matching is ambiguous, read candidate modules and choose by concrete file-pattern overlap.
- If multiple modules match, include all of them.
- If router defines always-load modules, include them.

3. Auto-select mode.
- If root `AGENTS.md` is missing, create one with `Overview`, `Base Policy Links (Load First)`, and `Local Details`.
- If root `AGENTS.md` exists, refresh `Base Policy Links (Load First)` to latest router/module links.
- Preserve valid repository-specific `Local Details`.
- Remove stale links that no longer match router or repository signals.

4. Write base policy links.
- Keep root content repository-specific.
- Add a link to `AGENTS.router.md`.
- Add links to selected module files.
- Use the latest default-branch links unless the user requested commit-pinned links.

5. Keep local details repository-specific.
- Keep concrete commands, folder ownership, runtime constraints, and explicit exceptions.
- Do not copy module policy text inline when links are sufficient.

6. Validate before finishing.
- Confirm each linked file exists on GitHub.
- Confirm module selection is justified by router conditions.
- Confirm no hardcoded technology mapping was introduced by the skill itself.
- Report exactly which remote files and branch or commit were used.

## Portability

- This `SKILL.md` is the canonical, provider-neutral instruction set.
- Adapter files for specific tools may exist, but they should mirror this file and must not weaken these rules.

## Universality Rules

- Re-learn from fresh `RevoTale/agent-docs` data on every run.
- Let router/module content drive technology selection instead of static stack lists.
- When new technology modules are added in `agent-docs`, they should be discovered and applied automatically if repository patterns match.
