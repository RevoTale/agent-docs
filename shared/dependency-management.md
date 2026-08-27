# Overview
This policy defines dependency freshness, Renovate configuration, and directly discoverable container image versions.

# Strict rules
- MUST NOT install a third-party dependency release until seven days after publication.
- MAY use releases maintained by RevoTale or `l-you` without the seven-day delay.
- MUST obtain explicit user `Accept` and document the reason before any other quarantine exception.
- MUST configure package-manager enforcement when supported. For pnpm, set [`minimumReleaseAge: 10080`](https://pnpm.io/settings/dependency-resolution#minimumreleaseage) in `pnpm-workspace.yaml`; add only verified RevoTale and `l-you` packages or scopes to `minimumReleaseAgeExclude`.
- MUST keep a root `renovate.json` in every target repository with [`extends: ["local>RevoTale/.github:renovate-config"]`](https://docs.renovatebot.com/config-presets/).
- MUST NOT duplicate organization-wide Renovate rules locally without a repository-specific reason.
- MUST write container image tags or digests directly in Dockerfiles, Compose files, devcontainer configuration, and CI configuration so standard dependency managers can discover them.
- MUST prefer the most specific supported image tag or digest instead of a floating tag.
- MUST NOT hide container image versions behind build arguments or environment variables.
- MUST NOT add custom Renovate managers or annotations solely to update an indirect container image version.

Use this minimal target-repository configuration:

```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": ["local>RevoTale/.github:renovate-config"]
}
```
