# Overview
This module defines authoring rules for repositories that use `templ` to render HTML from Go.
It covers rendering boundaries, component ownership, styling, composition, and client-side integration.
See paired examples in [../../examples/templ/index.md](../../examples/templ/index.md).
For deeper templ syntax and behavior details, see `https://templ.guide/llms.md`.

# Project structure
```text
<go-templ-repo-root>/
  AGENTS.md
  go.mod
  <feature>/
    <component>.templ
```

# Strict rules
- MUST read enforced utility/library choices from [../../awesome/index.md](../../awesome/index.md) and [../../awesome/go.md](../../awesome/go.md) before introducing, replacing, or removing templ usage.
  Example: [../../examples/templ/awesome-registry.md](../../examples/templ/awesome-registry.md)
- MUST use `github.com/a-h/templ` for the `html-templating` capability in Go repositories.
  Example: [../../examples/templ/templ-library-choice.md](../../examples/templ/templ-library-choice.md)
- MUST keep generated templ Go files in sync with `*.templ` sources.
  Example: [../../examples/templ/generated-files-sync.md](../../examples/templ/generated-files-sync.md)
- MUST keep templ components focused on rendering and light presentation logic.
  Example: [../../examples/templ/rendering-boundary.md](../../examples/templ/rendering-boundary.md)
- MUST prepare display-specific data in Go before render when domain models do not match the UI shape, and SHOULD prefer explicit view models at that boundary.
  Example: [../../examples/templ/view-model-boundary.md](../../examples/templ/view-model-boundary.md)
- MUST NOT call repositories, mailers, external APIs, or application services from templ components.
  Example: [../../examples/templ/no-service-calls.md](../../examples/templ/no-service-calls.md)
- SHOULD pass data through component parameters by default.
  Example: [../../examples/templ/explicit-parameters.md](../../examples/templ/explicit-parameters.md)
- SHOULD use the implicit `ctx` sparingly for cross-cutting request-scoped data such as locale, theme, or authenticated user context.
  Example: [../../examples/templ/ctx-usage.md](../../examples/templ/ctx-usage.md)
- MUST keep each visual component family self-contained in one owning `.templ` file when markup and local styling belong together.
  Example: [../../examples/templ/self-contained-component-family.md](../../examples/templ/self-contained-component-family.md)
- SHOULD define component-local selectors with `css name() { ... }` in the same `.templ` file as the component that owns them.
  Example: [../../examples/templ/local-css-ownership.md](../../examples/templ/local-css-ownership.md)
- MUST treat templ CSS components as generated classes and MUST NOT depend on their generated names across files or from JS.
  Example: [../../examples/templ/generated-class-opacity.md](../../examples/templ/generated-class-opacity.md)
- MUST apply templ CSS via `class={ ... }`, using CSS component calls, strings, maps, or `templ.KV(...)`, and MUST NOT hand-write generated hashed class names.
  Example: [../../examples/templ/class-expression-usage.md](../../examples/templ/class-expression-usage.md)
- MUST keep pseudo-elements, pseudo-classes, descendant or relational selectors, media queries, and keyframes in normal CSS rather than templ CSS components.
  Example: [../../examples/templ/complex-selectors-normal-css.md](../../examples/templ/complex-selectors-normal-css.md)
- MUST add stable raw class names when selectors or contracts need to cross component boundaries or be consumed by JS, Alpine, tests, or external CSS.
  Example: [../../examples/templ/stable-contract-classes.md](../../examples/templ/stable-contract-classes.md)
- SHOULD prefer templ CSS components over raw `<style>` blocks when the styling is local, flat, and owned by one component family.
  Example: [../../examples/templ/prefer-templ-css.md](../../examples/templ/prefer-templ-css.md)
- SHOULD keep one CSS component when the structure is the same and only runtime values such as color, size, offset, delay, width, or progress change.
  Example: [../../examples/templ/css-variable-ownership.md](../../examples/templ/css-variable-ownership.md)
- SHOULD prefer CSS custom properties for value-only variations over creating many near-duplicate CSS components.
  Example: [../../examples/templ/css-variable-ownership.md](../../examples/templ/css-variable-ownership.md)
- SHOULD keep static defaults in `css {}` or global CSS, and use `style={ ... }` mainly for per-render overrides from an owning parent or root.
  Example: [../../examples/templ/css-variable-ownership.md](../../examples/templ/css-variable-ownership.md)
- MUST prefer one component with parameters and conditional classes over duplicated markup branches when the structure is the same. 
  Example: [../../examples/templ/component-variants.md](../../examples/templ/component-variants.md)
- MUST use `children...` for one unnamed wrapped body.
  Example: [../../examples/templ/children-body.md](../../examples/templ/children-body.md)
- MUST use `templ.Component` parameters for explicit named slots such as `header`, `body`, `footer`, or `actions`.
  Example: [../../examples/templ/named-slots.md](../../examples/templ/named-slots.md)
- MUST use `templ.Attributes` spread for wrapper-style composition and pass-through attributes.
  Example: [../../examples/templ/attributes-spread.md](../../examples/templ/attributes-spread.md)
- SHOULD use standard `<script>` tags and standalone JS modules for client behavior that spans multiple elements or needs stable DOM hooks.
  Example: [../../examples/templ/standard-script-modules.md](../../examples/templ/standard-script-modules.md)
- MAY use `templ.JSFuncCall` or script templates for direct `on*` handlers when that is the simplest safe fit.
  Example: [../../examples/templ/js-func-call.md](../../examples/templ/js-func-call.md)
- MUST use stable hooks such as `data-*`, `id`, or `x-ref` for JS or Alpine targeting and MUST NOT query templ-generated CSS classes from client code.
  Example: [../../examples/templ/stable-hooks.md](../../examples/templ/stable-hooks.md)
- MUST declare `templ.NewOnceHandle()` at package scope when a component needs to emit scripts, styles, or dependencies once per page render context.
  Example: [../../examples/templ/once-handle.md](../../examples/templ/once-handle.md)
- SHOULD use `templ.JSONString` or `templ.JSONScript` when client code needs structured server data.
  Example: [../../examples/templ/structured-json.md](../../examples/templ/structured-json.md)

# Working Agreements
- MUST follow root interaction protocol from [../../AGENTS.md](../../AGENTS.md) before finalizing policy changes.
- MUST ask whether a selector is local component styling or a shared cross-component contract when that ownership is ambiguous.
- MUST ask whether `ctx` usage is intentional when the same data could be passed explicitly as component parameters.
