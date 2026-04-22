# Legacy Script Templates
Treat script templates as a legacy feature and prefer standard `<script>` tags, `templ.JSFuncCall`, `templ.JSONString`, and `templ.JSONScript` for new code.

## Pattern
```templ
package components

templ DashboardChart(data ChartData) {
  <section data-chart-root>
    @templ.JSONScript("chart-data", data)
    <script type="module" src="/assets/chart.js"></script>
  </section>
}
```

The pattern separates JavaScript loading from server data and uses the current templ helpers for the data handoff.

## Anti-pattern
```templ
package components

script initChart(data ChartData) {
  renderChart(data)
}

templ DashboardChart(data ChartData) {
  <section data-chart-root>
    @initChart(data)
  </section>
}
```

The anti-pattern achieves the same chart initialization, but it uses a legacy script template instead of the preferred new-code path.
