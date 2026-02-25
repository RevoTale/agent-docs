# Overview
This module defines baseline rules that apply to every repository regardless of stack.

# Project structure
```text
<repo-root>/
  AGENTS.md
```

# Strict rules
- Reuse Taskfile workflow rules for all projects: [../taskfile/doc.md](../taskfile/doc.md)
- Use Taskfile as the default workflows runner across modules.
- Keep this module scoped to repository-wide baseline rules only; stack-specific rules belong in stack modules.
- Add rules incrementally; each new rule must be concrete, testable, and scoped.

# Working Agreements
- Apply this module to every target repository.
- Keep this module free of stack-specific constraints.
