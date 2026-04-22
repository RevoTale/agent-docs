# Standard Script Modules
Use standard `<script>` tags and standalone JS modules for client behavior that spans multiple elements or needs stable DOM hooks.

## Pattern
```templ
package components

templ Dashboard() {
  <section data-dashboard>
    <button type="button" data-dashboard-refresh>Refresh</button>
    <div data-dashboard-status>Up to date</div>
  </section>
}

templ DashboardPage() {
  @Dashboard()
  <script type="module" src="/assets/dashboard.js"></script>
}
```

The pattern keeps the component focused on markup and loads reusable behavior through a normal JS module at the page or root boundary.

## Anti-pattern
```templ
package components

templ Dashboard() {
  <section data-dashboard>
    <button type="button" data-dashboard-refresh>Refresh</button>
    <div data-dashboard-status>Up to date</div>
    <script>
      const root = document.querySelector("[data-dashboard]");
      const button = root.querySelector("[data-dashboard-refresh]");
      const status = root.querySelector("[data-dashboard-status]");

      button.addEventListener("click", async () => {
        status.textContent = "Refreshing";
        await refreshDashboard();
        status.textContent = "Up to date";
      });
    </script>
  </section>
}

templ DashboardPage() {
  @Dashboard()
}
```

The anti-pattern achieves the same refresh behavior, but embeds reusable client logic inside the component instead of loading a standalone module.
