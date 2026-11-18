# Dependencies and Renovate

## Third-party release age

Wait seven days after publication before installing a third-party dependency release. Releases maintained by RevoTale or `l-you` may be used immediately. Any other exception requires explicit maintainer approval and a documented reason.

Configure enforcement when the package manager supports it. For pnpm, set [`minimumReleaseAge: 10080`](https://pnpm.io/settings/dependency-resolution#minimumreleaseage) and `minimumReleaseAgeIgnoreMissingTime: false` in `pnpm-workspace.yaml`. Limit `minimumReleaseAgeExclude` to verified RevoTale and `l-you` packages or scopes.

## Renovate preset

Keep the following configuration at the repository root in `renovate.json`. Extend the [organization preset](https://docs.renovatebot.com/config-presets/) locally only when a project-specific requirement calls for it.

```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": ["local>RevoTale/.github:renovate-config"]
}
```

## Container references

Write image tags or digests directly in Dockerfiles, Compose files, Dev Container configuration, and CI. Prefer the most specific supported tag with its digest, such as `<image>:1.2.3@sha256:<digest>`. Renovate can discover these references directly. Keep image versions out of build arguments and environment variables, and avoid custom managers or annotations solely to track indirect versions.
