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

## Nested AGENTS.md

Target repositories may use a root `AGENTS.md` plus nested `AGENTS.md` files when different subprojects need more specific stack rules.

Example:

```text
<repo-root>/
  AGENTS.md
  frontend/
    AGENTS.md
    package.json
  backend/
    AGENTS.md
    go.mod
```

Use nested files for meaningful architecture boundaries such as apps, services, or packages with distinct tooling. The nearest `AGENTS.md` should apply to the active subtree, while parent files stay additive for shared rules.

## Skill Workflows

Use the skill that matches the repository stage.

### Greenfield Init

Skill path: `skills/init-project-from-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/init-project-from-agent-docs skill to design the initial root and nested AGENTS.md files for this new repository.`

### Existing Repo AGENTS Refresh

Skill path: `skills/refresh-project-agents-from-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/refresh-project-agents-from-agent-docs skill to refresh the root and nested AGENTS.md files for this repository.`

### Existing Repo Refactor

Skill path: `skills/refactor-project-to-agent-docs/SKILL.md`

Example request:
`Use https://github.com/RevoTale/agent-docs/skills/refactor-project-to-agent-docs skill to align this repository and its nested subprojects with the recommended stack after an architecture interview.`

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
