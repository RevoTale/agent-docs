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
- Use Taskfile as the default interface for codegen, fix, validation, and test workflows across modules.

# Working Agreements
- Every project should follow the Taskfile conventions in [../taskfile/AGENTS.md](../taskfile/AGENTS.md).
- Add rules incrementally; each new rule must be concrete, testable, and scoped.
