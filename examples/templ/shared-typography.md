# Shared Typography
Use shared typography primitives or a global typography layer when the same typography treatment is repeated across multiple component families.

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
  color: var(--text-strong);
}

css footerTitle() {
  font-size: 0.95rem;
  color: var(--text-muted);
}

templ Card(title string) {
  <h3 class={ displayHeading(), cardTitle() }>{ title }</h3>
}

templ FooterSection(title string) {
  <h4 class={ displayHeading(), footerTitle() }>{ title }</h4>
}
```

The pattern extracts the repeated display-heading treatment once and composes it with local typography deltas.

## Anti-pattern
```templ
package components

css cardTitle() {
  font-family: var(--font-display);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 1.1rem;
  color: var(--text-strong);
}

css footerTitle() {
  font-family: var(--font-display);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 0.95rem;
  color: var(--text-muted);
}

templ Card(title string) {
  <h3 class={ cardTitle() }>{ title }</h3>
}

templ FooterSection(title string) {
  <h4 class={ footerTitle() }>{ title }</h4>
}
```

The anti-pattern renders the same titles, but duplicates the shared typography treatment in each local class.
