# Typography Local Deltas
Keep shared font family, weight, casing, and tracking in the shared definition, and keep per-component size, spacing, and color deltas local unless they are identical too.

## Pattern
```templ
package components

css displayHeading() {
  font-family: var(--font-display);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

css cardTitle() {
  font-size: 1.1rem;
  margin-block-end: 0.4rem;
  color: var(--text-strong);
}

css footerTitle() {
  font-size: 0.95rem;
  margin-block-end: 0.2rem;
  color: var(--text-muted);
}

templ Card(title string) {
  <h3 class={ displayHeading(), cardTitle() }>{ title }</h3>
}

templ FooterSection(title string) {
  <h4 class={ displayHeading(), footerTitle() }>{ title }</h4>
}
```

The pattern keeps the shared heading treatment reusable and leaves component-specific spacing, size, and color local.

## Anti-pattern
```templ
package components

css cardDisplayHeading() {
  font-family: var(--font-display);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 1.1rem;
  margin-block-end: 0.4rem;
  color: var(--text-strong);
}

css footerDisplayHeading() {
  font-family: var(--font-display);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 0.95rem;
  margin-block-end: 0.2rem;
  color: var(--text-muted);
}

templ Card(title string) {
  <h3 class={ cardDisplayHeading() }>{ title }</h3>
}

templ FooterSection(title string) {
  <h4 class={ footerDisplayHeading() }>{ title }</h4>
}
```

The anti-pattern drags local deltas into the shared definition, which forces multiple near-duplicate typography primitives instead of one reusable base plus local classes.
