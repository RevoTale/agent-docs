# Overview
This repository is the organization-wide AGENTS policy source.
It stores strict project rules and defines architecture policy including:
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
  shared/               # reusable cross-stack rulesets. extracted shared rules
    <rule-name>.md
  skills/               # directory containing skills that helps intergating/communicating with current rules
    <skill-name>.md
  modules/              # directory containing stack-specific guidance
    <module-name>/AGENTS.md
```

# Root AGENTS.md rules
- `AGENTS.md` documents this repository only and does not define downstream project rules directly.
- Root `AGENTS.md` should not update itself when new stack or shared rules are being added. Keep it universal and describe the general design and rules to follow only for the current project.

# Instructions

## Root doc.md
- Routing logic is maintained in `doc.md`.
- `doc.md` must include references to all baseline and stack modules.
- `doc.md` must provide a canonical table with short stack key, full stack name, module path, and `load_when`.
- `doc.md` must instruct agents to always load baseline modules and load project-specific stacks by `load_when` signals.
- Changes to module paths or routing signals must update `doc.md` in the same change.
- When adding a new stack, update `doc.md` with both the short stack key and full stack name.
- Root `doc.md` should contain only routing/composition logic and helpers to assemble target `AGENTS.md`.

## Stack
- Stack-specific guidance is maintained in `modules/*/AGENTS.md`.
- If equivalent rules are shared by multiple stack modules, extract them to `shared/<rule-name>.md`.
- Stack modules must link extracted shared rules by relative path (for module files: `[shared/<rule-name>.md](../../shared/<rule-name>.md)`).
- Keep shared files concrete and tool-focused; stack modules should keep only stack-specific additions.

## Rules combination
Rules in this project must be combined so a target agent can merge them into a single `AGENTS.md` using the contract defined in `doc.md`.

# Working Agreements

When user tells to add or update current docs:
  1. learn the target topic
  2. learn the current documentation
  3. learn the recommended ways to write clean `AGENTS.md` for both humans and agents. Be concise.
  4. learn what user actually asked for. If the task is ambiguous, ask questions, adapt to current documentation structure, and propose better alternatives when useful.
