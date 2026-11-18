# Maintain the RevoTale engineering standard

Start with [README.md](README.md) for the catalog, scope and precedence of project rules. These instructions cover contributions to the central repository.

## Change a rule

Put stack guidance in `modules/<stack>/doc.md`, shared rules in `shared/`, and library choices in the [registry](awesome/index.md). Link to the owner of a rule wherever another module needs it. Update the matching registry when adding, renaming or moving a document. Keep consumer links within the material a project can select.

Give each distinct decision a descriptive heading. A short rule can be a sentence. For longer decisions, use **Why**, **Rule** and **Result** when they clarify the problem, required action and observable outcome. Put comparisons in [paired examples](examples/pattern-anti-pattern.md). Use project checks for enforceable constraints. Keep maintainer procedures here.

## Change the workflow

The [adoption guide](examples/adoption/README.md) owns the sync procedure. The [skill](skills/apply-agent-docs/SKILL.md) guides architecture selection, adoption and migration. Keep engineering requirements in the documents linked from the catalog. Record agreed decisions in those documents and obtain maintainer acceptance before finalizing a policy change.

## Validate

Run the same checks as [Woodpecker](.woodpecker/validate-agent-docs.yaml) from the repository root.

```sh
go test ./scripts/validate-agent-docs.go ./scripts/validate-agent-docs_test.go
go run ./scripts/validate-agent-docs.go
```

For changes to selection, also exercise a new project and an existing project with a different runtime. Check the selected files and links and verify that unrelated tooling survives adoption.
