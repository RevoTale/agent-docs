# Rendering Boundary
Keep templ components focused on rendering and light presentation logic.

## Pattern
```templ
package components

css badgeRoot() {
  padding: 0.25rem 0.5rem;
}

css badgePaid() {
  border-color: var(--signal-strong);
}

css badgeError() {
  border-color: var(--signal-weak);
}

templ StatusBadge(status string, failed bool) {
  label := status
  if failed {
    label = "failed"
  }

  <span
    class={
      badgeRoot(),
      templ.KV(badgePaid(), status == "paid"),
      templ.KV(badgeError(), failed),
    }
  >
    { label }
  </span>
}
```

The pattern receives render-ready state and keeps the template focused on presentation decisions.

## Anti-pattern
```templ
package components

css badgeRoot() {
  padding: 0.25rem 0.5rem;
}

css badgePaid() {
  border-color: var(--signal-strong);
}

css badgeError() {
  border-color: var(--signal-weak);
}

templ StatusBadge(orderID string, svc OrderService) {
  order, err := svc.Load(ctx, orderID)
  if err != nil {
    <span class={ badgeRoot(), badgeError() }>failed</span>
    return
  }
  <span class={ badgeRoot(), templ.KV(badgePaid(), order.Status == "paid") }>{ order.Status }</span>
}
```

The anti-pattern renders the same badge, but crosses the render boundary by loading application data inside the template.
