# Overview
This module defines baseline rules that apply to every repository regardless of stack.

# Folder Structure
```text
<repo-root>/
  AGENTS.md
```

# Core Behaviors & Patterns
- Reuse Taskfile workflow rules for all projects: [../taskfile/AGENTS.md](../taskfile/AGENTS.md)

# Conventions
- Use Taskfile as the default workflows runner across modules.

- Keep this module scoped to repository-wide baseline rules only; stack-specific rules belong in stack modules.
- Add rules incrementally; each new rule must be concrete, testable, and scoped.
