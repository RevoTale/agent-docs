---
name: refactor-project-to-agent-docs
description: Interview-first workflow for aligning an existing repository with the current agent-docs architecture. Understand what the app does and what changes are acceptable, propose the target stack and refactor plan, wait for explicit Accept, then update AGENTS.md, tooling, and code toward compliance.
---

# Refactor Project To agent-docs Compliance

Use this workflow when the goal is to align an existing project with `agent-docs` policy modules and verify that alignment.

## Non-Negotiables

- Start with an architecture interview and repository audit.
- Show the target architecture and phased plan before any writes.
- Wait for explicit `Accept` on the target architecture and plan before mutating files.
- Use current `agent-docs` files as source of truth, not memory.
- Preserve behavior unless the user permits breaking changes.
- Keep edits incremental with clear checkpoints.
- End with an explicit compliance report.

## Workflow

1. Run the architecture interview.
- Capture app purpose, critical flows, runtime surfaces, data and storage, auth and identity, integrations, deployment constraints, mandated technologies, and allowed migration risk.
- Separate what may change now from what must stay stable.
- If the repository intent is already documented, confirm it instead of re-asking everything.

2. Load the current `agent-docs` source of truth.
- Resolve `doc.md` first, then `awesome/index.md`.
- Load always-on modules, matched stack modules, and matching awesome files.
- Use repository signals and interview answers together to choose the target architecture.
- Treat the current repository `AGENTS.md` as local input to preserve valid repository-specific exceptions, not as the router source of truth.

3. Audit the target project against policy.
- Inspect structure, dependency and runtime setup, task runner config, lint and format config, tests, build entrypoints, and major code patterns.
- Identify mismatches between current project state and loaded policy rules.
- Capture baseline command results for lint, test, and build when those commands exist.

4. Propose the target architecture and plan.
- Summarize selected modules, awesome capabilities, and required library or tooling changes.
- Output an ordered phase plan with goals, impacted files, risks, validations, and rollback notes for risky steps.
- Cover at minimum: structure, task runner/tasks, linting/formatting, code-style normalization, dependency alignment, and `AGENTS.md` updates.
- Ask for explicit `Accept` before editing.

5. Execute by phase.
- Refactor structure to the approved policy-aligned layout and naming.
- Refactor automation and tasking to match expected workflow patterns.
- Refactor lint, format, and dependency setup to the approved stack and capability choices.
- Refactor code style and architecture incrementally while preserving approved behavior.
- Update `AGENTS.md` and developer documentation so commands and rules remain accurate.

6. Verify `agent-docs` compliance.
- Build a rule-by-rule checklist from the loaded modules and awesome files.
- Mark each rule as `pass`, `fail`, or `not-applicable` with file or command evidence.
- Fix `fail` items that are in scope.

7. Validate the refactor.
- Run formatter, linter, tests, and build checks after changes.
- Report regressions fixed and any checks that could not run.

8. Deliver the final report.
- Summarize changes by phase.
- Provide the final `agent-docs` compliance checklist.
- List assumptions, residual gaps, and follow-ups.

## Default Decisions When Policy Is Silent

- Prefer one canonical task entrypoint per workflow (`format`, `lint`, `test`, `build`, `dev`).
- Prefer deterministic auto-fix style rules.
- Prefer minimal tool sprawl and explicit command ownership.

## Output Contract

Always return sections in this order:

1. Architecture Interview Summary
2. Refactor Plan
3. Applied Changes
4. agent-docs Compliance Report
5. Validation Results
6. Assumptions And Follow-Ups
