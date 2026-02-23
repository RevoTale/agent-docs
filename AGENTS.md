# Overview
This repository is the organization-wide AGENTS policy source. It stores shared stack modules and the router that selects which modules to load for a target repository.

# Folder Structure
Current repository layout:

```text
<repo-root>/
  AGENTS.md                                  # repository description (this file)
  AGENTS.router.md                           # module routing and load conditions
  shared/                                    # reusable cross-stack rulesets
    <rule-name>.md                           # extracted shared rules
  skills/
    init-project-from-agent-docs/
      SKILL.md         # init skill
  modules/ # directory containing stack-specific guidance

```

# Working Agreements
- `AGENTS.md` documents this repository only and does not define downstream project rules directly.
- Routing logic is maintained in `AGENTS.router.md`.
- Stack-specific guidance is maintained in `modules/*/AGENTS.md`.
- `AGENTS.router.md` must include references to all stack modules.
- Root `AGENTS.md` should not update itself when new stack or shared rules are being added. Keep it universal and describe the general design.
- Changes to module paths or routing signals must update `AGENTS.router.md` in the same change.
- When adding a new stack, update `AGENTS.router.md` with both the short stack key and full stack name.
- `AGENTS.router.md` must instruct agents to load only project-specific stack modules.
- If equivalent rules are shared by multiple stack modules, extract them to `shared/<rule-name>.md`.
- Stack modules must link extracted shared rules by relative path (for module files: `[/shared/<rule-name>.md](../../shared/<rule-name>.md)`).
- Keep shared files concrete and tool-focused; stack modules should keep only stack-specific additions.
- When user tells to add something/update in current docs:
  - learn the target topic
  - learn the current documentation
  - learn the reccomended ways to write clean `AGENTS.md` for both: humans and agents. Be concise.
  - learn what user actually ask for, if the target task is ambigious -> as the questions
  - adapt the received task to fit the current documentation structrue; propose the better alternatives if there are some