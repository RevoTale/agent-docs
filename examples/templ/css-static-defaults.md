# CSS Static Defaults
Keep static defaults in `css {}` or global CSS, and use `style={ ... }` mainly for per-render overrides from an owning parent or root.

## Pattern
```templ
package components

css pageBody() {
  --button-bg: var(--surface-strong);
  --button-fg: var(--text-strong);
}

css buttonRoot() {
  background: var(--button-bg, var(--surface-strong));
  color: var(--button-fg, var(--text-strong));
}

templ Button(label string) {
  <button class={ buttonRoot() }>{ label }</button>
}

templ DashboardActions(highlight bool) {
  if highlight {
    <section
      style={
        templ.KV("--button-bg", "var(--signal-strong)"),
        templ.KV("--button-fg", "white"),
      }
    >
      @Button("Save")
    </section>
    return
  }

  <section>
    @Button("Save")
  </section>
}

templ Page(highlight bool) {
  <body class={ pageBody() }>
    @DashboardActions(highlight)
  </body>
}
```

The pattern keeps static defaults in CSS and uses `style={ ... }` only for a real per-render override.

## Anti-pattern
```templ
package components

css buttonRoot() {
  background: var(--button-bg, var(--surface-strong));
  color: var(--button-fg, var(--text-strong));
}

templ Button(label string) {
  <button class={ buttonRoot() }>{ label }</button>
}

templ DashboardActions() {
  <section
    style={
      templ.KV("--button-bg", "var(--surface-strong)"),
      templ.KV("--button-fg", "var(--text-strong)"),
    }
  >
    @Button("Save")
  </section>
}

templ Page() {
  <body>
    @DashboardActions()
  </body>
}
```

The anti-pattern pushes static defaults into inline style output on every render instead of keeping them in CSS.
