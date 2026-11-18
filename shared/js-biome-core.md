# Biome checks

Use [Biome](https://biomejs.dev/) as the JavaScript-family linter and formatter. Configure `linter.rules.preset` as `all` to enable stable rules. Keep global nursery rules disabled because they fall outside Biome's semantic-versioning guarantees.

Make warnings fail validation with `biome check --error-on-warnings .` or its equivalent. Keep suppressions local and explain why the rule does not apply.
