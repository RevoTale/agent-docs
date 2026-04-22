# CSS Variable Ownership
Prefer CSS custom properties for value-only visual variations instead of creating many near-duplicate CSS components.

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

css buttonPrimary() {
  border-color: var(--signal-strong);
}

templ Button(label string, primary bool) {
  <button class={ buttonRoot(), templ.KV(buttonPrimary(), primary) }>
    { label }
  </button>
}

templ DashboardActions(primary bool) {
  sectionBg := "var(--surface-strong)"
  sectionFg := "var(--text-strong)"
  if primary {
    sectionBg = "var(--signal-strong)"
    sectionFg = "white"
  }

  <section
    style={
      templ.KV("--button-bg", sectionBg),
      templ.KV("--button-fg", sectionFg),
    }
  >
    @Button("Save", primary)
  </section>
}

templ Page(primary bool) {
  <body class={ pageBody() }>
    <main>
      @DashboardActions(primary)
    </main>
  </body>
}
```

The pattern keeps one button style and varies only the values through custom properties.

## Anti-pattern
```templ
package components

css buttonNeutral() {
  background: var(--surface-strong);
  color: var(--text-strong);
}

css buttonPrimaryTheme() {
  background: var(--signal-strong);
  color: white;
}

css buttonPrimary() {
  border-color: var(--signal-strong);
}

templ Button(label string, primary bool) {
  <button class={ buttonNeutral(), templ.KV(buttonPrimary(), primary), templ.KV(buttonPrimaryTheme(), primary) }>
    { label }
  </button>
}

templ DashboardActions(primary bool) {
  <section>
    @Button("Save", primary)
  </section>
}

templ Page(primary bool) {
  <body>
    <main>
      @DashboardActions(primary)
    </main>
  </body>
}
```

The anti-pattern duplicates theme-specific CSS classes instead of keeping one structure and varying only the values.
