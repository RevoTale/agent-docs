---
name: init-project-from-agent-docs
description: Interview-first workflow for new or near-empty repositories. Learn what the app should do, select the matching agent-docs modules and awesome capabilities, propose the initial architecture and AGENTS.md, and wait for explicit Accept before writing files.
---

# Initialize A Project From agent-docs

Use this skill when the repository is empty or early-stage and the user wants the initial architecture and `AGENTS.md` generated from product intent instead of existing code signals.

## Workflow

1. Run a short architecture interview.
- Ask only the questions needed to choose architecture: app purpose, user-facing surfaces, runtime shape (`web app`, `API`, `worker`, `CLI`), data and storage, auth and identity, key integrations, deployment constraints, and any mandated technologies.
- If an answer is missing but a safe default exists, state the assumption explicitly instead of blocking.
- If the repository already has substantial code or tooling, switch to `refresh-project-agents-from-agent-docs` or `refactor-project-to-agent-docs`.

2. Load the current policy source of truth.
- Resolve `doc.md` first, then `awesome/index.md`.
- Load always-on modules from `doc.md`.
- Load stack modules whose `load_when` conditions match the interview answers or any existing repository signals.
- Load matching awesome files by stack and capability.
- Do not hardcode technology names when the router or awesome registry can decide them.

3. Propose the initial architecture before writing.
- Summarize the project brief in 3-6 bullets.
- Show selected modules and why they match.
- Show selected awesome capabilities and the required libraries they imply.
- Draft the target `AGENTS.md` structure with `Overview`, `Base Policy Links (Load First)`, and `Local Details`.
- Call out assumptions, open questions, and non-default choices.

4. Wait for explicit approval.
- Ask for explicit `Accept` before creating or updating `AGENTS.md` or suggesting code scaffolding.
- If the user changes scope, revise the proposal and ask again.

5. Write the initial `AGENTS.md`.
- Create a repository-specific `Overview`.
- Add links to `doc.md`, `awesome/index.md`, always-on modules, and matched stack modules.
- Use latest default-branch links unless the user asked for commit-pinned links.
- Keep `Local Details` limited to project-specific constraints, decisions, and approved exceptions.
- Do not inline module policy text when links are sufficient.
- Do not copy awesome registry tables into the target `AGENTS.md`.

6. Validate and report.
- Confirm every linked file exists in the selected `agent-docs` revision.
- Report which router, module, and awesome files were selected and why.
- Report any assumptions that still need user confirmation.
