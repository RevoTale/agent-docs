# Stable Contract Classes
Add stable raw classes when a selector is a cross-component styling contract, and prefer `data-*`, `id`, or `x-ref` for JS hooks.

## Pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

templ Card(title string) {
  <article class={ cardRoot(), "card-root" }>
    <h3>{ title }</h3>
  </article>
}
```

```css
.dashboard-grid > .card-root {
  margin-block-start: 1rem;
}
```

The pattern adds a stable raw class only for the external CSS contract.

## Anti-pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

templ Card(title string) {
  <article class={ cardRoot() }>
    <h3>{ title }</h3>
  </article>
}
```

```css
.dashboard-grid > .templ_2f3a91 {
  margin-block-start: 1rem;
}
```

The anti-pattern relies on a generated class name for a cross-file CSS contract, which breaks as soon as templ regenerates it.
