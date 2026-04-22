# templ Library Choice
Use `github.com/a-h/templ` for Go HTML templating in repositories that match this module.

## Pattern
```templ
package components

templ Header(title string) {
  <h1>{ title }</h1>
}
```

## Anti-pattern
```go
package components

import "html/template"

var headerTemplate = template.Must(template.New("header").Parse(`<h1>{{ .Title }}</h1>`))
```

The anti-pattern switches to another HTML templating system instead of using the required `templ` capability.
