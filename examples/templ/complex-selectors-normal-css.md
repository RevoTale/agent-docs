# Complex Selectors In Normal CSS
Keep pseudo-selectors, descendant selectors, media queries, and keyframes in normal CSS.

templ CSS components do not support these patterns at all. The anti-pattern below is intentionally invalid.

## Pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

templ Card(title string) {
  <article class={ cardRoot(), "card-root" }>
    <h3 class="card-title">{ title }</h3>
  </article>
}
```

```css
.card-root:hover .card-title {
  color: var(--accent-strong);
}

@media (width >= 48rem) {
  .card-root {
    padding: 1.5rem;
  }
}
```

The pattern keeps flat local declarations in templ CSS and moves unsupported selector patterns to normal CSS.

## Anti-pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

css cardRoot:hover {
  color: var(--accent-strong);
}

css cardRoot() {
  @media (width >= 48rem) {
    padding: 1.5rem;
  }
}

templ Card(title string) {
  <article class={ cardRoot() }>
    <h3>{ title }</h3>
  </article>
}
```

The anti-pattern tries to put hover and media-query selectors into templ CSS, but templ CSS only supports flat class-owned declarations.
