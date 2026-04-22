# Stable Hooks
Use stable hooks such as `data-*`, `id`, or `x-ref` for JS or Alpine targeting and do not query templ-generated CSS classes from client code.

## Pattern
```templ
package components

css dropdownRoot() {
  display: grid;
  gap: 0.4rem;
}

templ Dropdown() {
  <div class={ dropdownRoot() } x-ref="root" data-dropdown>
    <button type="button" data-dropdown-trigger>Open</button>
  </div>
}
```

```js
const root = document.querySelector("[data-dropdown]");
const trigger = root?.querySelector("[data-dropdown-trigger]");
```

The pattern gives client code a stable contract that survives templ regeneration.

## Anti-pattern
```templ
package components

css dropdownRoot() {
  display: grid;
  gap: 0.4rem;
}

templ Dropdown() {
  <div class={ dropdownRoot() }>
    <button type="button">Open</button>
  </div>
}
```

```js
const root = document.querySelector(".templ_92ab11");
const trigger = root?.querySelector("button");
```

The anti-pattern keeps the same component shape, but omits stable hooks and forces client code to target an unstable generated class instead.
