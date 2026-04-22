# Local CSS Ownership
Define component-local selectors with `css name() { ... }` in the same `.templ` file as the component that owns them.

## Pattern
```templ
package components

css cardRoot() {
  display: grid;
  gap: 0.75rem;
  padding: 1rem;
}

css cardTitle() {
  font-weight: 700;
}

templ Card(title string) {
  <article class={ cardRoot() }>
    <h3 class={ cardTitle() }>{ title }</h3>
  </article>
}
```

## Anti-pattern
```templ
// styles.templ
package components

css cardRoot() {
  padding: 1rem;
}

// card.templ
package components

templ Card(title string) {
  <article class={ cardRoot() }>
    <h3>{ title }</h3>
  </article>
}
```

The anti-pattern hides a component-local selector in a different file even though only `Card` owns it.
