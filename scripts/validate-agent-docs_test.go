package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, relPath string, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, relPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	return path
}

func writeTempDoc(t *testing.T, content string) string {
	t.Helper()

	return writeTempFile(t, "doc.md", content)
}

func TestCheckNormativeBulletsPassesForMustBullets(t *testing.T) {
	path := writeTempDoc(t, `# Strict rules
- MUST keep this deterministic.
`)

	v := &validator{}
	checkNormativeBullets(v, path, sectionStrictRules)

	if len(v.failures) != 0 {
		t.Fatalf("expected no failures, got: %v", v.failures)
	}
}

func TestCheckNormativeBulletsFailsForNonNormativeBullets(t *testing.T) {
	path := writeTempDoc(t, `# Strict rules
- Keep this deterministic.
`)

	v := &validator{}
	checkNormativeBullets(v, path, sectionStrictRules)

	if len(v.failures) == 0 {
		t.Fatalf("expected at least one failure, got none")
	}
}

func TestCheckWorkingAgreementSemanticsRejectsTechnicalCommand(t *testing.T) {
	path := writeTempDoc(t, `# Working Agreements
- MUST follow root protocol.
- MUST run task validate before finalizing.
`)

	v := &validator{}
	checkWorkingAgreementSemantics(v, path)

	if len(v.failures) == 0 {
		t.Fatalf("expected technical command failure, got none")
	}
}

func TestParseAwesomeRegistryPaths(t *testing.T) {
	path := writeTempDoc(t, `# Awesome Registry
| stack_key | awesome_file | scope |
| --- | --- | --- |
| go | [go.md](./go.md) | go stack |
| react | [react.md](./react.md) | react stack |
`)

	paths, err := parseAwesomeRegistryPaths(path)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	if len(paths) != 2 {
		t.Fatalf("expected 2 awesome paths, got %d (%v)", len(paths), paths)
	}

	if paths[0] != filepath.Clean("awesome/go.md") {
		t.Fatalf("unexpected first path: %s", paths[0])
	}

	if paths[1] != filepath.Clean("awesome/react.md") {
		t.Fatalf("unexpected second path: %s", paths[1])
	}
}

func TestParseSkillDocParsesFrontmatter(t *testing.T) {
	path := writeTempFile(t, "skills/init-project-from-agent-docs/SKILL.md", `---
name: init-project-from-agent-docs
description: Initialize a repository from intent.
---

# Skill
Interview the user first.
`)

	doc, err := parseSkillDoc(path)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	if doc.name != "init-project-from-agent-docs" {
		t.Fatalf("unexpected skill name: %s", doc.name)
	}

	if doc.description != "Initialize a repository from intent." {
		t.Fatalf("unexpected skill description: %s", doc.description)
	}
}

func TestCheckSkillFileRejectsDeprecatedRouterReference(t *testing.T) {
	path := writeTempFile(t, "skills/refresh-project-agents-from-agent-docs/SKILL.md", `---
name: refresh-project-agents-from-agent-docs
description: Refresh AGENTS for an existing repository.
---

# Skill

Use repository signals and ask for Accept.
Load AGENTS.router.md before matching modules.
`)

	v := &validator{}
	checkSkillFile(v, path)

	if len(v.failures) == 0 {
		t.Fatalf("expected stale router reference failure, got none")
	}
}

func TestCheckSkillFileRequiresInterviewForInitSkill(t *testing.T) {
	path := writeTempFile(t, "skills/init-project-from-agent-docs/SKILL.md", `---
name: init-project-from-agent-docs
description: Initialize a repository from intent.
---

# Skill

Ask for explicit Accept before writing files.
`)

	v := &validator{}
	checkSkillFile(v, path)

	if len(v.failures) == 0 {
		t.Fatalf("expected interview requirement failure, got none")
	}
}
