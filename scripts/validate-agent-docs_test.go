package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func put(t *testing.T, root, name, body string) {
	t.Helper()
	name = filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func fixture(t *testing.T, snapshot bool) string {
	t.Helper()
	root := t.TempDir()
	put(t, root, "modules/common/doc.md", "# Common\nUse [Taskfile](../taskfile/doc.md#commands) and [OS defaults][os].\n[os]: ../../shared/os.md\n")
	put(t, root, "modules/taskfile/doc.md", "# Taskfile\nUse the [tool](https://taskfile.dev).\n")
	put(t, root, "shared/os.md", "# OS\n```md\n[Literal example](missing.md)\n```\n")
	put(t, root, "LICENSE", "An upstream license.\n")
	if !snapshot {
		put(t, root, "README.md", "| [common](modules/common/doc.md) | Always |\n| [taskfile](modules/taskfile/doc.md) | Always |\n")
		put(t, root, "awesome/index.md", "| auth | [Auth](./auth.md) | Identity |\n")
		put(t, root, "awesome/auth.md", "| capability | status | package |\n| --- | --- | --- |\n| identity | required | kratos |\n")
		put(t, root, skillPath, "---\nname: apply-agent-docs\ndescription: Select and update the project standard.\n---\nRead the catalog.\n")
	}
	return root
}

func TestCompleteStandardAndSelectedSnapshot(t *testing.T) {
	for _, snapshot := range []bool{false, true} {
		if got := validate(fixture(t, snapshot), snapshot); len(got) != 0 {
			t.Errorf("snapshot=%t: %v", snapshot, got)
		}
	}
}

func TestValidatingSnapshotErrors(t *testing.T) {
	cases := []struct {
		name    string
		mutate  func(*testing.T, string)
		wantErr string
	}{
		{"missing reference definition destination", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "shared/os.md")); err != nil {
				t.Fatal(err)
			}
		}, "missing consumer document"},
		{"outside root", func(t *testing.T, root string) {
			put(t, root, "shared/os.md", "[read](../../AGENTS.md)")
		}, "missing consumer document"},
		{"unsupported URL", func(t *testing.T, root string) {
			put(t, root, "shared/os.md", "[read](file:///etc/hosts)")
		}, "expected a relative policy link"},
		{"symlink directory", func(t *testing.T, root string) {
			target := t.TempDir()
			put(t, target, "hidden.md", "# Hidden\n")
			if err := os.Symlink(target, filepath.Join(root, "shared/alias")); err != nil {
				t.Fatal(err)
			}
			put(t, root, "shared/os.md", "[read](alias/hidden.md)")
		}, "regular, non-executable file"},
		{"executable markdown", func(t *testing.T, root string) {
			if err := os.Chmod(filepath.Join(root, "shared/os.md"), 0o755); err != nil {
				t.Fatal(err)
			}
		}, "regular, non-executable file"},
		{"central author instructions", func(t *testing.T, root string) {
			put(t, root, "AGENTS.md", "# Maintainers\n")
		}, "outside the consumer documentation layout"},
		{"nested skill", func(t *testing.T, root string) {
			put(t, root, "shared/nested/SKILL.md", "# Skill\n")
		}, "outside the consumer documentation layout"},
		{"missing baseline", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "modules/common/doc.md")); err != nil {
				t.Fatal(err)
			}
		}, "baseline module required"},
		{"missing license", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "LICENSE")); err != nil {
				t.Fatal(err)
			}
		}, "include the upstream license"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, true)
			tc.mutate(t, root)
			if got := strings.Join(validate(root, true), "\n"); !strings.Contains(got, tc.wantErr) {
				t.Errorf("expected %q, got %q", tc.wantErr, got)
			}
		})
	}
}

func TestInvalidCatalog(t *testing.T) {
	root := fixture(t, false)
	put(t, root, "modules/react/doc.md", "# React\n")
	put(t, root, "awesome/react.md", "| CSS | required | Tailwind |\n| testing | default | testtool |\n")
	put(t, root, "awesome/index.md", "| one | ./auth.md |\n| two | ./auth.md |\n")
	put(t, root, skillPath, "---\nname: old-name\ndescription: Select a snapshot.\n---\nA workflow.\n")
	got := strings.Join(validate(root, false), "\n")
	for _, problem := range []string{"module missing from catalog", "library list missing from catalog", "expected one entry", "invalid library status", "matching name"} {
		if !strings.Contains(got, problem) {
			t.Errorf("expected %q, got %q", problem, got)
		}
	}
}
