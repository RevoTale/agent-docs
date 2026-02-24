# Overview
This repository is the organization-wide AGENTS policy source.
It stores the strict to rules for the projects and defines strict policies enforcing good architecture pattern including:
- technologies which to choose
- code style
- code quality tools
- naming conventions
- directory structure
- code pattern
- custom strict rules

Based on the target project, techonology stack should be combined into single `AGENTS.md` at the root of target project via appropriate `SKILL`. It assembles all project-specific stack modules keeping all the convention and rules

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
    <module-name>.md
```

# Root AGENTS.md rules
- `AGENTS.md` documents this repository only and does not define downstream project rules directly.
- Root `AGENTS.md` should not update itself when new stack or shared rules are being added. Keep it universal and describe the general design and rules to follow only for the current project.

# Instructions

## Root doc.md
- Routing logic is maintained in `doc.md`.
- `doc.md` must include references to all stack modules.
- `doc.md` must instruct agents to load only project-specific stack modules.
- Changes to module paths or routing signals must update `doc.md` in the same change.
- When adding a new stack, update `doc.md` with both the short stack key and full stack name.
- Root `doc.md` should contain only logic only a logic and helpers to assemble the target `AGENTS.md`.

## Stack
- Stack-specific guidance is maintained in `modules/*/AGENTS.md`.
- If equivalent rules are shared by multiple stack modules, extract them to `shared/<rule-name>.md`.
- Stack modules must link extracted shared rules by relative path (for module files: `[shared/<rule-name>.md](../../shared/<rule-name>.md)`).
- Keep shared files concrete and tool-focused; stack modules should keep only stack-specific additions.

## Rules combination
All the rules in this project should be combined in the way, the target AGENT can merge them easily into a single `AGENTS.md` via one of the techniques defined in the `doc.md`.

# Working Agreements

When user tells to add something/update in current docs:
  1. learn the target topic
  2. learn the current documentation
  3. learn the reccomended ways to write clean `AGENTS.md` for both: humans and agents. Be concise.
  4. learn what user actually ask for. If the target task is ambigious -> ask the questions -> adapt the received information about task to fit the current documentation structrue; propose the better alternatives if there are some