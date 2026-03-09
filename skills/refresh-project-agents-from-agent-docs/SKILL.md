---
name: refresh-project-agents-from-agent-docs
description: Refresh or create AGENTS.md for an existing repository by reading concrete repository signals, selecting the matching agent-docs modules and awesome files, proposing the AGENTS diff, and waiting for explicit Accept before writing.
---

# Refresh AGENTS.md From agent-docs

Use this skill when the codebase already exists and the goal is to create or sync `AGENTS.md` without refactoring the application itself.

## Workflow

1. Load the current `agent-docs` source of truth.
- Resolve `doc.md` first, then `awesome/index.md`.
- Load always-on modules from `doc.md`.
- Fetch only the module docs and awesome files needed for matching and validation.
- Never use deprecated router aliases; resolve the canonical router file from `doc.md`.

2. Detect repository signals.
- Inspect concrete signals such as manifests, lockfiles, source extensions, framework folders, task runners, lint configs, and runtime or deploy configs.
- Evaluate `load_when` conditions from `doc.md` against those signals.
- If multiple modules match, include all of them.
- If the repository is empty or signals are too weak to infer intent, switch to `init-project-from-agent-docs` and interview the user.
- If the user requests code or tooling changes beyond `AGENTS.md`, switch to `refactor-project-to-agent-docs`.

3. Propose the `AGENTS.md` selection.
- List the repository signals that justified each selected module.
- List selected awesome files by stack and capability.
- Show the intended `AGENTS.md` changes, including removed stale links.
- Preserve a repository-specific `Overview` and valid `Local Details`.

4. Wait for explicit approval.
- Ask for explicit `Accept` before writing or rewriting `AGENTS.md`.

5. Apply the refresh.
- Create `AGENTS.md` if it is missing.
- Update `Base Policy Links (Load First)` to the selected current policy links.
- Use latest default-branch links unless the user asked for commit-pinned links.
- Preserve project-specific constraints, commands, ownership notes, and approved exceptions in `Local Details`.
- Remove stale policy links that no longer match the repository or the current router.

6. Validate and report.
- Confirm every linked file exists in the selected `agent-docs` revision.
- Report the selected modules, awesome files, and repository signals used to choose them.
- Report assumptions or ambiguities that still need user confirmation.
