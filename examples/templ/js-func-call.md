# JSFuncCall
Use `templ.JSFuncCall` for small direct `on*` handlers when that is the simplest safe fit.

## Pattern
```templ
package components

templ InlineAction(message string) {
  <button onClick={ templ.JSFuncCall("console.log", message) }>
    Debug
  </button>
}
```

The pattern uses templ's JS helper to serialize the handler call safely.

## Anti-pattern
```templ
package components

templ InlineAction(message string) {
  <button onClick={ "console.log('" + message + "')" }>
    Debug
  </button>
}
```

The anti-pattern hand-builds event JavaScript as a raw string instead of using templ's JS helpers.
