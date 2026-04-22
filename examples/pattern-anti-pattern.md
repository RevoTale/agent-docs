# Pattern And Anti-pattern Examples

This document defines how example files in this repository should teach rules to humans and AI systems.

Examples in this repo are comparison artifacts.
They are designed to help a reader or another AI infer the intended rule with low ambiguity and apply it correctly in the requested task.

Examples are not primarily tutorials.
They should optimize for clear comparison, not broad coverage.

# Core Idea

A strong example file should behave like a controlled comparison:

- same task
- same context
- same rough artifact shape when reasonable
- one important difference
- clear consequence

The purpose is to make the intended rule easy to infer.

If a pattern and anti-pattern differ in many unrelated ways, the reader or model cannot reliably determine which difference matters.

# Scope

This guidance applies to examples for any artifact type, including:

- code
- configuration
- templates
- API contracts
- architecture docs
- workflow docs
- policy docs
- source vs generated file boundaries

# Authoring Rules

- MUST let each example file teach one rule only.
- MUST include both `Pattern` and `Anti-pattern` in the same file.
- MUST make the pattern and anti-pattern solve the same task.
- MUST keep the same rough inputs, context, and artifact shape when reasonable.
- MUST isolate one main difference between the two sides.
- MUST make the anti-pattern plausible unless the rule is specifically about invalid usage.
- MUST make the consequence of the anti-pattern visible from the example or the short explanation.
- MUST keep examples aligned with the official or enforced source of truth for the technology or policy being taught.
- SHOULD prefer short, retrieval-friendly examples over broad tutorial-style examples.
- SHOULD place the preferred `Pattern` before the `Anti-pattern`.
- SHOULD keep explanations short and use them only to clarify the decision boundary.
- SHOULD avoid adding extra abstractions or side details unrelated to the rule being taught.

# Required File Shape

Use this structure by default:

````md
# <Rule Name>

Short sentence explaining what the rule means.

## Pattern

```<language-or-format>
...
```

Short sentence explaining why this is preferred.

## Anti-pattern

```<language-or-format>
...
```

Short sentence explaining why this is worse.
````

Use the most natural fence or format for the artifact:
- `go`
- `templ`
- `ts`
- `json`
- `yaml`
- `md`
- `text`

# Comparison Design

## What should stay the same

Keep these stable between the pattern and anti-pattern whenever reasonable:

- task
- audience or usage context
- inputs
- output intent
- rough structure
- artifact type

## What should change

Change only the part that teaches the rule, for example:

- ownership boundary
- data flow
- naming
- dependency choice
- selector contract
- API shape
- configuration placement
- generation workflow
- source-of-truth choice
- documentation structure
- architectural boundary
- safety or validation rule

# Good Comparison Properties

A good comparison has these properties:

- same task, different decision
- realistic mistake, not nonsense
- one dominant lesson
- visible failure mode
- low noise
- easy retrieval

# Anti-pattern Guidance

Anti-patterns should represent mistakes that a reasonable engineer or AI might actually produce.

Good anti-patterns often show:

- wrong ownership
- hidden coupling
- unstable contracts
- duplicated structure for the same job
- misuse of global or ambient context
- mixing layers of responsibility
- violating a source-of-truth boundary
- fragile workflow ordering
- unclear or unsafe configuration
- documentation that teaches the wrong contract

Do not use anti-patterns that are weak only because:

- they are random nonsense
- they contain obvious syntax errors unless invalid syntax is the point
- they solve a different task than the pattern
- they silently violate several unrelated rules at once
- they depend on hidden context not shown in the file

# Validity

When possible, both the pattern and anti-pattern should be valid and plausible.

Why:
- realistic comparisons teach better than broken ones
- valid anti-patterns show actual decision mistakes
- AI learns better from believable alternatives than from nonsense

Use invalid examples only when the rule is specifically about invalid syntax, forbidden constructs, or generator errors.

# When Exact Symmetry Is Not Possible

Sometimes the pattern and anti-pattern cannot be structurally identical.

In that case:

- keep the same task first
- keep the same context where possible
- minimize structural drift
- explain the decision boundary in one short sentence

"Same task" matters more than "same exact code."

# Non-code Examples

This guidance also applies to non-code artifacts.

## Documentation examples

Use the same topic and audience on both sides.
Change only the documentation decision being taught.

Examples:
- clear contract vs ambiguous contract
- source-of-truth reference vs duplicated stale explanation
- concise actionable runbook vs broad narrative with missing steps

## Workflow examples

Use the same workflow step on both sides.
Change only the decision or ordering that matters.

Examples:
- edit source then regenerate vs edit generated output directly
- validate before merge vs skip validation
- read enforced dependency registry vs choose tools ad hoc

## Policy examples

Use the same policy intent on both sides.
Change only the rule structure or ownership boundary being taught.

Examples:
- technical invariant in `Strict rules` vs interaction rule in the wrong section
- stack-specific rule in the module vs cross-stack rule duplicated across modules

# Review Checklist

Before finalizing an example, verify:

- Does this file teach exactly one rule?
- Do pattern and anti-pattern solve the same task?
- Do they keep the same rough shape?
- Is the anti-pattern realistic?
- Is the main difference easy to spot?
- Is the consequence clear?
- Does the pattern avoid contradicting other repo rules?
- Is the example aligned with the official source of truth?
- Would another AI likely infer the intended lesson from this file alone?

# Design Principle

Optimize examples for comparison, not coverage.

A smaller, sharper example pair is better than a broader example that mixes several lessons at once.
