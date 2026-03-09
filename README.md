# RevoTale Agent Docs

Central source for shared AGENTS policies of [RevoTale](https://revotale.com).

## Source of Truth

- `AGENTS.md` and `doc.md` are the only policy sources of truth.
- `skills/` contains operational workflows that must consume those sources of truth rather than redefine them.
- `README.md` is an informative mirror and must stay aligned with those files.

## Repository Layout

- `AGENTS.md`: repository-level policy contract.
- `doc.md`: canonical module registry, routing logic, and merge contract.
- `awesome/index.md`: entrypoint for enforced utility/library choices by stack or capability.
- `awesome/<name>.md`: stack-specific or capability-specific enforced utility/library lists.
- `shared/<rule-name>.md`: reusable cross-stack rulesets referenced by modules.
- `modules/common/doc.md`: always-loaded baseline module.
- `modules/<stack>/doc.md`: stack-specific modules.
- `skills/<skill-name>/SKILL.md`: operational skills for greenfield init, AGENTS refresh, and guided refactors.

## Skill Workflows

Use the skill that matches the repository stage.

### Greenfield Init

Skill path: `skills/init-project-from-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/init-project-from-agent-docs skill to design the initial AGENTS.md for this new repository.`

### Existing Repo AGENTS Refresh

Skill path: `skills/refresh-project-agents-from-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/refresh-project-agents-from-agent-docs skill to refresh AGENTS.md for this repository.`

### Existing Repo Refactor

Skill path: `skills/refactor-project-to-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/refactor-project-to-agent-docs skill to align this repository with the recommended stack after an architecture interview.`

## Update Manually

You can still update `AGENTS.md` manually when automation is not desired.

1. Create a root `AGENTS.md` if missing.
2. Add `https://github.com/RevoTale/agent-docs/blob/main/doc.md` under `Base Policy Links (Load First)`.
3. Add `https://github.com/RevoTale/agent-docs/blob/main/modules/common/doc.md` under `Base Policy Links (Load First)`.
4. Add `https://github.com/RevoTale/agent-docs/blob/main/modules/taskfile/doc.md` under `Base Policy Links (Load First)`.
5. Add `https://github.com/RevoTale/agent-docs/blob/main/awesome/index.md` under `Base Policy Links (Load First)`.
6. Add matching stack module links from `modules/<stack>/doc.md` based on router conditions.
7. Add repository-specific rules under `Local Details`.

## Root `AGENTS.md` Example

```md
# Overview
Payments service API.

# Base Policy Links (Load First)
- https://github.com/RevoTale/agent-docs/blob/main/doc.md
- https://github.com/RevoTale/agent-docs/blob/main/modules/common/doc.md
- https://github.com/RevoTale/agent-docs/blob/main/modules/taskfile/doc.md
- https://github.com/RevoTale/agent-docs/blob/main/awesome/index.md

# Local Details
- Add repository-specific constraints and local working agreements.
```
