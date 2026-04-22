# Component Variants
Prefer one component with parameters and conditional classes over duplicated markup branches when the structure is the same.

## Pattern
```templ
package components

css buttonRoot() {
  border: 1px solid var(--line-soft);
  padding: 0.75rem 1rem;
}

css buttonPrimary() {
  border-color: var(--signal-strong);
}

templ Button(label string, primary bool) {
  <button class={ buttonRoot(), templ.KV(buttonPrimary(), primary) }>
    { label }
  </button>
}

templ DashboardActions() {
  <div>
    @Button("Save", true)
    @Button("Cancel", false)
  </div>
}
```

## Anti-pattern
```templ
package components

css buttonRoot() {
  border: 1px solid var(--line-soft);
  padding: 0.75rem 1rem;
}

css buttonPrimary() {
  border-color: var(--signal-strong);
}

templ PrimaryButton(label string) {
  <button class={ buttonRoot(), buttonPrimary() }>
    { label }
  </button>
}

templ SecondaryButton(label string) {
  <button class={ buttonRoot() }>
    { label }
  </button>
}

templ DashboardActions() {
  <div>
    @PrimaryButton("Save")
    @SecondaryButton("Cancel")
  </div>
}
```

The anti-pattern keeps the same markup shape but duplicates the component boundary instead of varying one component with parameters and conditional classes.
