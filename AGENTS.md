# Overview
This repository is the organization-wide AGENTS policy source.
It stores strict project rules that define architecture policy including:
- technology choices
- code style
- code quality tools
- naming conventions
- directory structure
- code patterns
- custom strict rules

Based on the target project, stack modules are combined into a single root `AGENTS.md` via the appropriate skill. It assembles project-specific rules while preserving shared conventions.

# Folder Structure
Current repository layout:

```text
<repo-root>/
  AGENTS.md             # repository description (this file)
  doc.md                # module routing and load conditions
  shared/               # reusable cross-stack rulesets
    <rule-name>.md
  skills/               # skills used to integrate with these policies
    <skill-name>.md
  modules/              # stack-specific guidance
    <module-name>/doc.md
```

# Root AGENTS.md rules
- `AGENTS.md` documents this repository only and does not define downstream project rules directly.
- Root `AGENTS.md` should not update itself when new stack or shared rules are added. Keep it universal and focused on repository-level design.

# Instructions

## Root doc.md
- Routing logic is maintained in `doc.md`.
- `doc.md` must include references to all baseline and stack modules.
- `doc.md` must provide a canonical table with short stack key, full stack name, module path, and `load_when`.
- `doc.md` must instruct agents to always load baseline modules and load project-specific stacks by `load_when` signals.
- `doc.md` must define section semantics: `Strict rules` for technical constraints, and `Working Agreements` for user-agent interaction protocol.
- Changes to module paths or routing signals must update `doc.md` in the same change.
- When adding a new stack, update `doc.md` with both the short stack key and full stack name.
- Root `doc.md` should contain only routing/composition logic and helpers to assemble target `AGENTS.md`.

## Stack
- Stack-specific guidance is maintained in `modules/*/doc.md`.
- If equivalent rules are shared by multiple stack modules, extract them to `shared/<rule-name>.md`.
- Stack modules must link extracted shared rules by relative path (for module files: `[shared/<rule-name>.md](../../shared/<rule-name>.md)`).
- Keep shared files concrete and tool-focused; stack modules should keep only stack-specific additions.

## Rules combination
Rules in this project must be combined so a target agent can merge them into a single `AGENTS.md` using the contract defined in `doc.md`.

# Working Agreements
- MUST ask interview questions when policy intent or requested changes are ambiguous.
- MUST wait for explicit `Accept` before finalizing any policy text change.
- MUST document user-approved technical exceptions in the relevant `Strict rules` section.
- MUST keep `Working Agreements` interaction-only; technical commands, checks, and invariants MUST be defined in `Strict rules`.
