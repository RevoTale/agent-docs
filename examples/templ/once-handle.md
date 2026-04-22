# Once Handle
Declare `templ.NewOnceHandle()` at package scope when a component needs to emit scripts, styles, or dependencies once per page render context.

## Pattern
```templ
package components

import "github.com/a-h/templ"

var chartDeps = templ.NewOnceHandle()

templ ChartDeps() {
  @chartDeps.Once() {
    <script type="module" src="/assets/chart.js"></script>
  }
}
```

## Anti-pattern
```templ
package components

import "github.com/a-h/templ"

templ ChartDeps() {
  deps := templ.NewOnceHandle()
  @deps.Once() {
    <script type="module" src="/assets/chart.js"></script>
  }
}
```

The anti-pattern recreates the handle inside the render path, which defeats the one-time contract.
