# children... Body
Use `children...` as the preferred API for one unnamed wrapped body.

This rule is about API shape. The alternative below is valid, but heavier than needed for a simple wrapper.

## Pattern
```templ
package components

css panelShell() {
  display: grid;
  gap: 0.8rem;
}

templ PanelShell() {
  <section class={ panelShell() }>
    { children... }
  </section>
}

templ DashboardBody() {
  <div>
    <h2>Overview</h2>
    <p>Body content</p>
  </div>
}

templ Dashboard() {
  @PanelShell() {
    @DashboardBody()
  }
}
```

The pattern gives a simple wrapper the natural wrapped-body call shape.

## Anti-pattern
```templ
package components

import "github.com/a-h/templ"

css panelShell() {
  display: grid;
  gap: 0.8rem;
}

templ PanelShell(body templ.Component) {
  <section class={ panelShell() }>
    @body
  </section>
}

templ DashboardBody() {
  <div>
    <h2>Overview</h2>
    <p>Body content</p>
  </div>
}

templ Dashboard() {
  @PanelShell(DashboardBody())
}
```

The anti-pattern still works, but it turns a simple unnamed wrapper into a heavier slot-style API with no added clarity.
