# Choosing a Go HTML renderer

A service needs to render HTML in Go.

## Pattern

Read the [Go library choices](../../awesome/go.md) and select the required `github.com/a-h/templ` package for the `html-templating` capability.

## Anti-pattern

Add `html/template` to implement the same capability. Checking the approved renderer after building the templates makes the switch more expensive.
