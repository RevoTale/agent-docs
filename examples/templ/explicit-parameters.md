# Explicit Parameters
Pass render data through component parameters by default.

## Pattern
```templ
templ UserSummary(name string, plan string, projectCount int) {
  <section>
    <h2>{ name }</h2>
    <p>{ plan }</p>
    <p>{ fmt.Sprintf("%d projects", projectCount) }</p>
  </section>
}
```

## Anti-pattern
```templ
type DashboardPageData struct {
  CurrentUser  User
  Projects     []Project
  FeatureFlags Flags
}

templ UserSummary(page DashboardPageData) {
  <section>
    <h2>{ page.CurrentUser.Name }</h2>
    <p>{ page.CurrentUser.Plan }</p>
    <p>{ fmt.Sprintf("%d projects", len(page.Projects)) }</p>
  </section>
}
```

Both versions render the same summary. The anti-pattern passes a page-sized data bucket even though the component only needs three render values.
