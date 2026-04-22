# Self-Contained Component Family
Keep one visual component family in one owning `.templ` file when its markup and local styling belong together.

## Pattern
```templ
// card.templ
package components

css cardRoot() {
  display: grid;
  gap: 0.75rem;
}

css cardTitle() {
  font-weight: 700;
}

templ Card(title string, body string) {
  <article class={ cardRoot() }>
    <h3 class={ cardTitle() }>{ title }</h3>
    <p>{ body }</p>
  </article>
}
```

The pattern keeps one local component family together because the markup and styling are owned by the same boundary.

## Anti-pattern
```templ
// card.templ
package components

templ Card(title string, body string) {
  @CardFrame() {
    @CardHeader(title)
    @CardBody(body)
  }
}

// card_frame.templ
package components

css cardRoot() {
  display: grid;
  gap: 0.75rem;
}

templ CardFrame() {
  <article class={ cardRoot() }>
    { children... }
  </article>
}

// card_header.templ
package components

css cardTitle() {
  font-weight: 700;
}

templ CardHeader(title string) {
  <h3 class={ cardTitle() }>{ title }</h3>
}

// card_body.templ
package components

templ CardBody(body string) {
  <p>{ body }</p>
}
```

The anti-pattern renders the same card, but scatters one local component family across files without creating any real reuse or ownership boundary.
