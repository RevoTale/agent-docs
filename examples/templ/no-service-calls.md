# No Service Calls
Do not call repositories, mailers, external APIs, or application services from templ components.

## Pattern
```go
func LoadInviteCard(ctx context.Context, userID string, svc InviteService) templ.Component {
  invites, err := svc.Load(ctx, userID)
  vm := NewInviteCardViewModel(invites, err)
  return InviteCard(vm)
}
```

```templ
templ InviteCard(vm InviteCardViewModel) {
  if vm.ErrorMessage != "" {
    <p>{ vm.ErrorMessage }</p>
    return
  }
  <div>{ vm.InviteCountLabel }</div>
}
```

## Anti-pattern
```templ
templ InviteCard(ctx context.Context, userID string, svc InviteService) {
  invites, err := svc.Load(ctx, userID)
  if err != nil {
    <p>Failed to load invites</p>
    return
  }
  <div>{ fmt.Sprintf("%d invites", len(invites)) }</div>
}
```

Both versions render the same invite card. The anti-pattern crosses the render boundary by performing the service call inside the template.
