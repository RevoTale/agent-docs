# Overview
This module defines baseline rules that apply to every repository regardless of stack.

# Project structure
```text
<repo-root>/
  AGENTS.md
```

# Strict rules
- MUST reuse Taskfile workflow rules for all projects: [../taskfile/doc.md](../taskfile/doc.md).
- MUST use Taskfile as the default workflow runner across modules.
- MUST keep this module scoped to repository-wide baseline rules only; stack-specific rules belong in stack modules.
- MUST require `task validate` to pass before merge.
- MUST require `task test` to pass when the project defines `task test`.
- SHOULD add rules incrementally, and each new rule SHOULD be concrete, testable, and scoped.

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask user to move stack-specific requests into the matching stack module instead of common.
