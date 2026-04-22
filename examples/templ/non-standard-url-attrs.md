# Non-standard URL Attributes
Use `templ.URL(...)` for dynamic URL values in non-standard URL attributes such as `hx-get`, `hx-post`, and similar client-library attributes.

## Pattern
```templ
package components

templ ContactRow(id string, name string) {
  <div hx-get={ templ.URL("/contacts/" + id + "/details") }>
    { name }
  </div>
}
```

The pattern uses templ's URL helper for a dynamic URL in a non-standard attribute.

## Anti-pattern
```templ
package components

templ ContactRow(id string, name string) {
  <div hx-get={ "/contacts/" + id + "/details" }>
    { name }
  </div>
}
```

The anti-pattern renders the same attribute, but it skips templ's URL sanitization path for non-standard URL attributes.
