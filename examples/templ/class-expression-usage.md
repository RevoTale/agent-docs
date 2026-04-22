# Class Expression Usage
Apply templ CSS through `class={ ... }` and `templ.KV(...)`, not by hand-writing generated class names.

## Pattern
```templ
package components

css badgeRoot() {
  padding: 0.25rem 0.5rem;
}

css badgePaid() {
  border-color: var(--signal-strong);
}

templ Badge(status string) {
  <span class={ badgeRoot(), templ.KV(badgePaid(), status == "paid") }>{ status }</span>
}
```

## Anti-pattern
```templ
package components

templ Badge(status string) {
  className := "templ_2f3a91"
  if status == "paid" {
    className += " templ_91bf2a"
  }
  <span class={ className }>{ status }</span>
}
```

The anti-pattern bypasses templ's class expression model and hard-codes unstable generated names.
