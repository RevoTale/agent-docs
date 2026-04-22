# View-Model Boundary
Prepare display-specific data in Go before rendering when domain data does not already match the UI shape.

## Pattern
```go
type InviteCardViewModel struct {
  InviteCountLabel string
  ErrorMessage     string
}

func NewInviteCardViewModel(invites []Invite, err error) InviteCardViewModel {
  if err != nil {
    return InviteCardViewModel{
      ErrorMessage: "Failed to load invites",
    }
  }

  return InviteCardViewModel{
    InviteCountLabel: fmt.Sprintf("%d invites", len(invites)),
  }
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

The pattern keeps formatting and error mapping in Go so the template receives one consistent render contract.

## Anti-pattern
```templ
templ InviteCard(user User, invites []Invite, err error) {
  totalLabel := fmt.Sprintf("%s has %d pending invites", user.DisplayName(), len(invites))
  if err != nil {
    if errors.Is(err, ErrBackendTimeout) {
      <p>Backend timeout while loading invites</p>
      return
    }
    <p>Unknown invite error</p>
    return
  }
  <div>{ totalLabel }</div>
}
```

The anti-pattern pushes formatting and error-mapping decisions into the template instead of shaping render-ready data first.
