---
name: apply-agent-docs
description: Adopt, update or migrate a project to the RevoTale engineering standard. Use when selecting modules for a new project, replacing an earlier AGENTS.md bootstrap or aligning an existing architecture.
---

# Apply the RevoTale standard

Use the [module catalog](../../README.md) and [adoption guide](../../examples/adoption/README.md) from the approved RevoTale revision. When this skill is installed separately, open those files in the trusted checkout used for the update. Engineering rules live in the selected documentation.

## Select a target

For a new project, use product requirements to identify deployable boundaries, runtimes and capabilities. Ask only about unresolved choices that would change the architecture.

For an existing project, inspect its manifests, lockfiles, application boundaries, checks and existing instructions. Identify the target and any gaps. A documentation refresh preserves the current runtime, local working agreements and scoped exceptions. For a requested migration, agree on the affected boundaries, compatibility requirements and first independently verifiable step.

## Apply the selection

Show the proposed selection and any architectural change requiring a decision. After approval, follow the adoption guide to sync a pinned copy and link the relevant files from the root or nested `AGENTS.md`. Preserve unrelated project instructions. An update includes the selected files, revision and any bootstrap adjustments in the same review.

For a migration, change the agreed scope and run the affected project checks after each independently useful step. Report outstanding gaps and evidence for the completed step. For ordinary development, the project's local bootstrap and checks are sufficient.
