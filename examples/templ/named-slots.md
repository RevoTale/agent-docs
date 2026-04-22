# Named Slots
Use `templ.Component` parameters for explicit named slots such as `header`, `body`, `footer`, or `actions`.

## Pattern
```templ
package components

import "github.com/a-h/templ"

templ BoardHeader(title string) {
  <h2>{ title }</h2>
}

templ BoardBody(summary string) {
  <p>{ summary }</p>
}

templ BoardActions() {
  <button type="button">Refresh</button>
}

templ Board(header templ.Component, body templ.Component, actions templ.Component) {
  <section>
    <header>@header</header>
    <div>@body</div>
    <footer>@actions</footer>
  </section>
}

templ Dashboard() {
  @Board(BoardHeader("Overview"), BoardBody("Builds are healthy"), BoardActions())
}
```

The pattern keeps slot meaning explicit in both the signature and the call site.

## Anti-pattern
```templ
package components

import "github.com/a-h/templ"

templ BoardHeader(title string) {
  <h2>{ title }</h2>
}

templ BoardBody(summary string) {
  <p>{ summary }</p>
}

templ BoardActions() {
  <button type="button">Refresh</button>
}

templ Board(children ...templ.Component) {
  <section>
    { children... }
  </section>
}

templ Dashboard() {
  @Board(BoardHeader("Overview"), BoardBody("Builds are healthy"), BoardActions())
}
```

The anti-pattern renders the same board, but hides slot meaning inside an unnamed child list.
