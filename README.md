# RevoTale Agent Docs

Central source for shared AGENTS policies of [RevoTale](https://revotale.com).

## Source of Truth

- `AGENTS.md` and `doc.md` are the only policy sources of truth.
- `README.md` is an informative mirror and must stay aligned with those files.

## Repository Layout

- `AGENTS.md`: repository-level policy contract.
- `doc.md`: canonical module registry, routing logic, and merge contract.
- `shared/<rule-name>.md`: reusable cross-stack rulesets referenced by modules.
- `modules/common/doc.md`: always-loaded baseline module.
- `modules/<stack>/doc.md`: stack-specific modules.
- `skills/init-project-from-agent-docs/`: skill to initialize or refresh target `AGENTS.md`.

## Install Or Update In A Target Repository

You can apply `agent-docs` in two ways: use the skill for automated create/update, or update `AGENTS.md` manually.

### Use The Skill

Skill path: `skills/init-project-from-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/init-project-from-agent-docs skill to refresh AGENTS.md for this repository.`

### Update Manually

1. Create a root `AGENTS.md` if missing.
2. Add `https://github.com/RevoTale/agent-docs/blob/main/doc.md` under `Base Policy Links (Load First)`.
3. Add `https://github.com/RevoTale/agent-docs/blob/main/modules/common/doc.md` under `Base Policy Links (Load First)`.
4. Add `https://github.com/RevoTale/agent-docs/blob/main/modules/taskfile/doc.md` under `Base Policy Links (Load First)`.
5. Add matching stack module links from `modules/<stack>/doc.md` based on router conditions.
6. Add repository-specific rules under `Local Details`.

## Root `AGENTS.md` Example

```md
# Overview
Payments service API.

# Base Policy Links (Load First)
- https://github.com/RevoTale/agent-docs/blob/main/doc.md
- https://github.com/RevoTale/agent-docs/blob/main/modules/common/doc.md
- https://github.com/RevoTale/agent-docs/blob/main/modules/taskfile/doc.md

# Local Details
- Add repository-specific constraints and local working agreements.
```
