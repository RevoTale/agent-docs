# This is a rules I created during the development of the root RevoTlae websire.

# Overview
`templ-rules.md` defines technology-level rules for writing `templ` components.

# Strict rules
- MUST make each visual component self-contained in its `.templ` file.
- MUST define component-local selectors with `css name() { ... }` in the same `.templ` file as the component that owns them.
- MUST treat templ `css name() { ... }` as a single generated class with a flat list of CSS properties. Use it for class-owned declarations only.
  Example: `css cardRoot() { padding: 1rem; }` is valid. `css cardRoot:hover { ... }`, `css cardRoot() { @media ... }`, `.parent:hover .child`, and `::before` are not valid templ `css` patterns.
- MUST apply internal selectors via `class={ rootClass(), partClass() }`. Never write generated hashed class names by hand.
- MUST treat templ `css` output as opaque. Generated class names are unstable and must not be depended on across files.
- MUST keep JS and Alpine targeting stable selectors during the refactor. Hashed templ classes remain styling-only.
- MUST prefer one owning `.templ` file per visual component family.
- MUST prefer templ `css` blocks over raw `<style>` blocks for component-local styling.
- MUST keep pseudo-elements, pseudo-classes, descendant or relational selectors, media queries, and keyframes out of templ `css {}`.
  Example: keep `css cardRoot() { padding: 1rem; }` in templ, but write `::before`, `:hover`, `.card-root:hover .card-title`, `@media`, and `@keyframes` in normal CSS.
- MUST add stable global class names when a component needs more complex selectors than templ `css {}` supports.
  Example:
  ```templ
  css cardRoot() {
    padding: 1rem;
  }

  templ Card() {
    <article class={ cardRoot(), "card-root" }>
      <h3 class="card-title">Hello</h3>
    </article>
  }
  ```
  Then keep selectors such as `.card-root::before`, `.card-root:hover .card-title`, `@media`, and `@keyframes` in normal CSS.
- MUST treat `templ.ComponentCSSClass` and `ClassName()` as advanced escape hatches, not the default pattern. Not recommended to use.
- MUST use stable raw class names only for cross-component contracts or non-CSS consumers.
  Example: `class={ cardRoot(), "surface-panel", "js-card" }` is valid when a shared style or JS selector must stay stable. `class="card-title"` is not valid when the title is only styled inside one component.
- MUST use `data-*`, `id`, `x-ref`, or explicit stable hook classes when JS or Alpine needs to target an element after render. Never query hashed templ classes from JS.
- MUST prefer component parameters, `templ.KV(...)`, and CSS custom properties over duplicated markup branches for variants.
  Example: prefer `Button(label, primary)` with `templ.KV(buttonPrimary(), primary)` over separate `PrimaryButton()` and `SecondaryButton()` components when the markup is the same.
- MUST prefer `children...`, `templ.Component` parameters, and `templ.Attributes` spread for composition.
- MUST use standard `<script>` tags, standalone JS modules, `templ.JSONString`, and `templ.JSFuncCall` for new JS integration work.
  Example: load `/assets/app.js` with `<script type="module" src="/assets/app.js"></script>` and keep `Alpine.data("dropdown", ...)` there instead of embedding ad-hoc inline scripts in each component.
- MUST NOT use legacy templ script templates for new code.
- MUST gate one-time script or dependency output with `templ.NewOnceHandle()` when a component needs to emit it once per page.
- SHOULD put runtime values such as color, size, offset, delay, or progress into CSS custom properties on the component root before creating many dynamic `css` variants.
  Example:
  ```templ
  css meterRoot() {
    inline-size: var(--meter-progress);
    background: var(--meter-accent);
  }

  templ Meter() {
    <div class={ meterRoot() }></div>
  }

  templ DashboardCard() {
    <section
      style={
        // If these values are static for this component, keep them in `css {}` instead.
        templ.KV("--meter-progress", "72%"),
        templ.KV("--meter-accent", "var(--signal-strong)"),
      }
    >
      <h3>Build progress</h3>
      @Meter()
    </section>
  }
  ```



# Primary pattern: isolated component
```templ
package components

css cardRoot() {
  display: grid;
  gap: 0.6rem;
  padding: 1rem;
  border: 1px solid var(--line-soft);
  background: var(--panel-bg);
}

css cardTitle() {
  font-family: var(--font-display);
  font-size: 1.1rem;
}

css cardUrgent() {
  border-color: var(--signal-strong);
}

templ Card(title string, urgent bool) {
  <article class={ cardRoot(), templ.KV(cardUrgent(), urgent) }>
    <h3 class={ cardTitle() }>{ title }</h3>
  </article>
}
```

# Pattern: wrapper with children
```templ
package components

import "github.com/a-h/templ"

css panelShell() {
  display: grid;
  gap: 0.8rem;
}

templ PanelShell(attrs templ.Attributes) {
  <section class={ panelShell() } { attrs... }>
    { children... }
  </section>
}

templ DockSection() {
  @PanelShell(templ.Attributes{"data-section": "dock"}) {
    <h2>Dock</h2>
    <p>Component body</p>
  }
}
```

# Pattern: named slots with `templ.Component`
```templ
package components

import "github.com/a-h/templ"

css boardRoot() {
  display: grid;
  gap: 1rem;
}

templ Board(header templ.Component, body templ.Component) {
  <section class={ boardRoot() }>
    <header>@header</header>
    <div>@body</div>
  </section>
}
```

Use `children...` for one unnamed wrapped body. Use `templ.Component` parameters when the slots are explicit, for example `header`, `body`, `footer`, or `actions`.

# Pattern: one-time script include
```templ
package components

import "github.com/a-h/templ"

var meterScript = templ.NewOnceHandle()

templ MeterDeps() {
  @meterScript.Once() {
    <script type="module" src="/assets/meter.js"></script>
  }
}
```

# Pattern: Alpine or JS hooks
```templ
package components

css dropdownRoot() {
  display: grid;
  gap: 0.4rem;
}

css dropdownPanel() {
  display: grid;
}

templ Dropdown() {
  <div class={ dropdownRoot() } x-data="dropdown()" x-ref="root" data-dropdown>
    <button type="button" x-on:click="open = !open" data-dropdown-trigger>
      Open
    </button>
    <div class={ dropdownPanel() } x-show="open" x-ref="panel" data-dropdown-panel>
      Contents
    </div>
  </div>
}
```

# Use These Replacements
- When JS or Alpine needs to find an element after render, use `data-*`, `id`, `x-ref`, or a stable hook class. Do not target templ-generated hashed classes.
- When a visual variant keeps the same structure, keep one component and vary it with parameters, `templ.KV(...)`, or CSS custom properties. Do not duplicate the whole markup tree.
- When adding client behavior, load normal JS modules with `<script>` and connect behavior through Alpine attributes or stable hooks. Do not default to anonymous inline scripts.
- When emitting one-time dependencies, declare `templ.NewOnceHandle()` once at package scope and reuse it. Do not create it inside a render path or component body.
 
# Source socs
To learn/find the missing features/syntax/docs read the `https://templ.guide/llms.md` file, which is designed specifically for the AI agents by the GO Templ maintainer.

# References
- https://github.com/a-h/templ/blob/main/docs/docs/03-syntax-and-usage/12-css-style-management.md
- https://github.com/a-h/templ/blob/main/docs/docs/03-syntax-and-usage/10-template-composition.md
- https://github.com/a-h/templ/blob/main/docs/docs/03-syntax-and-usage/13-script-templates.md
- https://github.com/a-h/templ/blob/main/docs/docs/03-syntax-and-usage/18-render-once.md