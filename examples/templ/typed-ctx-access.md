# Typed ctx Access
Read request-scoped context through small typed helper functions rather than raw `ctx.Value(...).(T)` expressions in markup.

## Pattern
```go
type themeKey struct{}

func ThemeFromContext(ctx context.Context) string {
  theme, _ := ctx.Value(themeKey{}).(string)
  return theme
}
```

```templ
package components

templ ThemeBadge() {
  theme := ThemeFromContext(ctx)
  <span data-theme={ theme }>{ theme }</span>
}
```

The pattern centralizes the context lookup behind a small typed helper and keeps markup readable.

## Anti-pattern
```go
type themeKey struct{}
```

```templ
package components

templ ThemeBadge() {
  <span data-theme={ ctx.Value(themeKey{}).(string) }>{ ctx.Value(themeKey{}).(string) }</span>
}
```

The anti-pattern repeats raw context lookups in markup and risks runtime panics when the key or type is wrong.
