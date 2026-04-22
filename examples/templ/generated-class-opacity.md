# Generated Class Opacity
Treat templ CSS component output as opaque and unstable across files and JS.

## Pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

templ Card() {
  <article class={ cardRoot() } data-card>
    Hello
  </article>
}
```

```js
const card = document.querySelector("[data-card]");
```

## Anti-pattern
```templ
package components

css cardRoot() {
  padding: 1rem;
}

templ Card() {
  <article class={ cardRoot() }>
    Hello
  </article>
}
```

```js
const card = document.querySelector(".templ_2f3a91");
```

Both versions try to find the same card element. The anti-pattern binds external code to a generated implementation detail instead of a stable contract.
