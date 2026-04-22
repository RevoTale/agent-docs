# Single CSS Component
Keep one CSS component when the structure is the same and only runtime values such as color, size, width, or progress change.

## Pattern
```templ
package components

css meterFill() {
  inline-size: var(--meter-width);
  background: var(--meter-color);
}

templ Meter(width string, color string) {
  <div
    style={
      templ.KV("--meter-width", width),
      templ.KV("--meter-color", color),
    }
  >
    <div class={ meterFill() }></div>
  </div>
}
```

The pattern keeps one CSS component and varies only the runtime values that feed it.

## Anti-pattern
```templ
package components

css meterStrong72() {
  inline-size: 72%;
  background: var(--signal-strong);
}

css meterMuted48() {
  inline-size: 48%;
  background: var(--surface-strong);
}

templ Meter(width string, color string) {
  className := meterMuted48()
  if width == "72%" && color == "var(--signal-strong)" {
    className = meterStrong72()
  }
  <div class={ className }></div>
}
```

The anti-pattern renders the same meter, but turns value combinations into separate CSS components that multiply as new states appear.
