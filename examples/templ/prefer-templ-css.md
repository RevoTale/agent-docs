# Prefer templ CSS
Prefer templ CSS components over raw `<style>` blocks when the styling is local, flat, and owned by one component family.

## Pattern
```templ
package components

css cardRoot() {
  display: grid;
  gap: 0.75rem;
  padding: 1rem;
}

templ Card(title string) {
  <article class={ cardRoot() }>
    <h3>{ title }</h3>
  </article>
}
```

## Anti-pattern
```templ
package components

templ Card(title string) {
  <style>
    .card-root {
      display: grid;
      gap: 0.75rem;
      padding: 1rem;
    }
  </style>
  <article class="card-root">
    <h3>{ title }</h3>
  </article>
}
```

The anti-pattern uses a raw `<style>` block for simple local styling that templ CSS can own more directly.
