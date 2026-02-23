---
name: refactor-project-to-agent-docs
description: Refactor the current repository in one prompt so its structure, tooling, and code conventions follow the active `agent-docs` policies, then verify compliance. Use when the user asks to adapt a project to `agent-docs` and expects plan-first execution with validation.
---

# Refactor Project To agent-docs Compliance

Use this workflow when the goal is to align an existing project with `agent-docs` policy modules and verify that alignment.

## Non-Negotiables

- Create and show the plan first.
- Execute immediately after the plan unless the user requests plan-only mode.
- Use current `agent-docs` files as source of truth, not memory.
- Preserve behavior unless the user permits breaking changes.
- Keep edits incremental with clear checkpoints.
- End with an explicit compliance report.

## Workflow

1. Load the active `agent-docs` baseline.
- Resolve policy source in this exact order:
  1) Fetch `https://raw.githubusercontent.com/RevoTale/agent-docs/main/AGENTS.router.md`
  2) After router is loaded, resolve module/shared files in the same order
- Load always-on modules and stack modules selected by router signals in the target project.
- Read linked `shared/*.md` rules used by selected modules.
- Respect precedence: target project local `AGENTS.md` overrides router and module defaults.

2. Audit the target project against policy.
- Inspect structure, dependency/runtime setup, task runner config, lint/format config, and key code-style patterns.
- Identify mismatches between current project state and loaded policy rules.
- Capture baseline command results for lint, test, and build (or nearest equivalents).

3. Create the plan first.
- Output an ordered phase plan with goals, impacted files, risks, and validations per phase.
- Cover at minimum: structure, task runner/tasks, linting/formatting, code-style normalization, and policy/documentation alignment.
- Include rollback guidance for high-risk edits.

4. Execute by phase.
- Refactor structure to policy-aligned layout and naming.
- Refactor automation/tasking to match expected `Taskfile` (or declared runner) patterns.
- Refactor lint/format setup to policy-required tools and settings.
- Refactor code style to align with policy while preserving behavior.
- Update AGENTS-related and developer docs so commands/rules are accurate.

5. Verify `agent-docs` compliance.
- Build a rule-by-rule checklist from the loaded modules.
- Mark each rule as `pass`, `fail`, or `not-applicable` with file/command evidence.
- Fix `fail` items that are in scope.

6. Validate the refactor.
- Run formatter, linter, tests, and build checks after changes.
- Report regressions fixed and any checks that could not run.

7. Deliver final report.
- Summarize changes by phase.
- Provide the final `agent-docs` compliance checklist.
- List assumptions, residual gaps, and follow-ups.

## Default Decisions When Policy Is Silent

- Prefer one canonical task entrypoint per workflow (`format`, `lint`, `test`, `build`, `dev`).
- Prefer deterministic auto-fix style rules.
- Prefer minimal tool sprawl and explicit command ownership.

## Output Contract

Always return sections in this order:

1. Refactor Plan
2. Applied Changes
3. agent-docs Compliance Report
4. Validation Results
5. Assumptions And Follow-Ups
