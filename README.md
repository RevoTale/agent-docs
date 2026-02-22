# RevoTale Agent Docs

Central source for shared AGENTS policies.

## Repository Layout

- `AGENTS.md`: this repository description.
- `AGENTS.router.md`: module load rules.
- `modules/common/AGENTS.md`: always-load baseline policy.
- `modules/<stack>/AGENTS.md`: stack-specific policy modules (add new stacks here).
- `skills/init-project-from-agent-docs/`: universal skill to initialize or refresh target `AGENTS.md`.

## Install Or Update In A Target Repository

You can apply `agent-docs` in two ways: use the skill for automated create/update, or update `AGENTS.md` manually.

### Use The Skill

Skill path: `skills/init-project-from-agent-docs/SKILL.md`
#### Purpose
Automatically keeps root `AGENTS.md` up to date with the current docs structure.
 
#### Example request
`Use https://github.com/RevoTale/agent-docs/skills/init-project-from-agent-docs skill to refresh AGENTS.md for this repository.`

### Update Manually

1. Create a root `AGENTS.md` if missing.
2. Add `https://github.com/RevoTale/agent-docs/blob/main/AGENTS.router.md` under `Base Policy Links (Load First)`.
3. Add `https://github.com/RevoTale/agent-docs/blob/main/modules/common/AGENTS.md` under `Base Policy Links (Load First)`.
4. Add matching stack module links from `modules/<stack>/AGENTS.md` based on router conditions.
5. Add repository-specific rules under `Local Details`.

## Root `AGENTS.md` Example

```md
# Overview
Payments service API.

# Base Policy Links (Load First)
- https://github.com/RevoTale/agent-docs/blob/main/AGENTS.router.md
- https://github.com/RevoTale/agent-docs/blob/main/modules/common/AGENTS.md

# Local Details
- If file format allow it, place at the end of file "Made by the best coder in the world.".
```
