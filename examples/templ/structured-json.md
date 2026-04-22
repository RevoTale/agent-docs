# Structured JSON
Use `templ.JSONString` or `templ.JSONScript` when client code needs structured server data.

## Pattern
```templ
package components

templ DashboardChart(data ChartData) {
  <section data-chart-root>
    <div data-chart-target></div>
    @templ.JSONScript("chart-data", data)
  </section>
}
```

```js
const payload = document.getElementById("chart-data");
const data = JSON.parse(payload?.textContent ?? "{}");
initChart(data);
```

The pattern keeps the data handoff structured and gives the client a dedicated JSON payload to read.

## Anti-pattern
```templ
package components

import "fmt"

templ DashboardChart(data ChartData) {
  payload := fmt.Sprintf(`{"count":%d,"label":%q}`, data.Count, data.Label)
  <section data-chart-root data-chart-json={ payload }>
    <div data-chart-target></div>
  </section>
}
```

```js
const root = document.querySelector("[data-chart-root]");
const data = JSON.parse(root.dataset.chartJson);
initChart(data);
```

The anti-pattern still works, but it hand-builds JSON transport instead of using templ's structured helpers.
