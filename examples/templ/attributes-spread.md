# Attributes Spread
Use `templ.Attributes` spread for wrapper-style composition and pass-through attributes.

## Pattern
```templ
package components

import "github.com/a-h/templ"

css panelShell() {
  display: grid;
  gap: 0.8rem;
}

templ PanelShell(attrs templ.Attributes) {
  <section class={ panelShell() } { attrs... }>
    { children... }
  </section>
}

templ Dashboard() {
  @PanelShell(templ.Attributes{
    "id": "overview",
    "aria-label": "Overview",
    "data-section": "overview",
  }) {
    <h2>Overview</h2>
    <p>Ready</p>
  }
}
```

The pattern keeps the wrapper API open for future attributes without growing the component signature.

## Anti-pattern
```templ
package components

templ PanelShell(id string, ariaLabel string, dataSection string) {
  <section
    class="panel-shell"
    id={ id }
    aria-label={ ariaLabel }
    data-section={ dataSection }
  >
    { children... }
  </section>
}

templ Dashboard() {
  @PanelShell("overview", "Overview", "overview") {
    <h2>Overview</h2>
    <p>Ready</p>
  }
}
```

The anti-pattern renders the same shell, but turns generic attribute passthrough into a brittle parameter list that grows over time.
