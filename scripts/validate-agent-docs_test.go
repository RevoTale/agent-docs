package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempDoc(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "doc.md")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp doc: %v", err)
	}

	return path
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
