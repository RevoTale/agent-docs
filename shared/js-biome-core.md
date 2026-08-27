# Overview
Reusable Biome baseline for JavaScript-family stacks.

# Core Behaviors & Patterns
- MUST use Biome as the linter and formatter: https://biomejs.dev/.
- MUST set `linter.rules.preset` to `all` so every stable rule is enabled.
- MUST NOT enable nursery rules globally because they are not covered by Biome semantic-versioning guarantees.
- MUST make Biome warnings fail validation, for example with `biome check --error-on-warnings .`.
- MUST keep suppressions local and explain why the rule does not apply.
- MUST NOT use ESLint or Prettier.
