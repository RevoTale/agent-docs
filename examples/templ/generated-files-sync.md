# Generated Files Sync
Keep generated templ Go files in sync with `*.templ` sources.

## Pattern
```text
components/button.templ
components/button_templ.go

task gen
```

## Anti-pattern
```text
1. Edit `components/button.templ`.
2. Do not run generation.
3. Commit stale `button_templ.go` output or edit it by hand.
```

The pattern treats generated Go files as outputs of the `.templ` source. The anti-pattern leaves generated code stale or manually modified.
