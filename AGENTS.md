# Overview
This repository contains the organization's universal AGENTS documentation system. It is a policy and routing repository that stores shared modules (common, Go, Next.js) and the router definition used to select them.

# Folder Structure
Current repository layout:

```text
<repo-root>/
  AGENTS.md            # repository description (this file)
  AGENTS.router.md     # module routing and load conditions
  skills/
    init-project-from-agent-docs/
      SKILL.md         # init skill
  modules/ # directory containg stack-specific guidance

```

# Working Agreements
- `AGENTS.md` documents this repository only and does not define downstream project rules directly.
- Routing logic is maintained in `AGENTS.router.md`.
- Stack-specific guidance is maintained in `modules/*/AGENTS.md`.
- Cross-agent initialization workflow is maintained in `skills/init-project-from-agent-docs/`.
- Changes to module paths or routing signals must update `AGENTS.router.md` in the same change.
- After we add a new stack, update the `AGENTS.router.md` to include it, with the short description and full name of the stack.
- `AGENTS.router.md` includes all references to the modules. 
- `AGENTS.router.md` should tell agents to load only project-specific stacks 