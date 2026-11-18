# Next.js

Apply the [React rules](../react/doc.md) and the [Next.js library choices](../../awesome/nextjs.md) for matching capabilities.

## Routing

Use the App Router in `app/` and keep framework configuration in `next.config.ts`. Introducing a Pages Router route requires explicit maintainer approval.

## Development and validation

Expose `next dev`, `next build`, and `next start` through the selected package manager. Taskfile can wrap these commands. For a Next.js application, `task validate` runs Biome checks only. This scope takes precedence over the general TypeScript validation workflow.
