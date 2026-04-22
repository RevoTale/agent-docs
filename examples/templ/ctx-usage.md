# ctx Usage
Use the implicit `ctx` for cross-cutting request-scoped data, not as a general data transport mechanism.

## Pattern
```go
type localeKey struct{}

func LocaleFromContext(ctx context.Context) string {
  locale, _ := ctx.Value(localeKey{}).(string)
  return locale
}
```

```templ
package components

templ DashboardHeader(userName string) {
  <header data-locale={ LocaleFromContext(ctx) }>
    <h1>Dashboard</h1>
    <p>{ userName }</p>
  </header>
}
```

The pattern keeps ordinary render data explicit and uses `ctx` only for a cross-cutting locale concern.

## Anti-pattern
```go
type localeKey struct{}
type userNameKey struct{}

func LocaleFromContext(ctx context.Context) string {
  locale, _ := ctx.Value(localeKey{}).(string)
  return locale
}

func UserNameFromContext(ctx context.Context) string {
  name, _ := ctx.Value(userNameKey{}).(string)
  return name
}
```

```templ
package components

templ DashboardHeader() {
  <header data-locale={ LocaleFromContext(ctx) }>
    <h1>Dashboard</h1>
    <p>{ UserNameFromContext(ctx) }</p>
  </header>
}
```

The anti-pattern renders the same header, but hides ordinary page data in ambient context instead of passing it as a parameter.
