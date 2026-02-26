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

