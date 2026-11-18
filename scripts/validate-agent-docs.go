package main

import (
	"flag"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

const skillPath = "skills/apply-agent-docs/SKILL.md"

var (
	moduleRow = regexp.MustCompile(`^\|\s*\[([^]]+)\]\((modules/[^)]+)\)\s*\|`)
	// Policy references use inline Markdown links or single-line reference definitions.
	inlineLink = regexp.MustCompile(`\]\(\s*(?:<([^>]+)>|([^\s)]+))`)
	definition = regexp.MustCompile(`^\s{0,3}\[[^]]+\]:\s*(?:<([^>]+)>|(\S+))`)
)

type checker struct {
	root     string
	docs     map[string]string
	failures []string
}

func (c *checker) fail(name, message string, args ...any) {
	c.failures = append(c.failures, name+": "+fmt.Sprintf(message, args...))
}

func (c *checker) read(name string) string {
	data, err := os.ReadFile(filepath.Join(c.root, filepath.FromSlash(name)))
	if err != nil {
		c.fail(name, "%v", err)
		return ""
	}
	if !utf8.Valid(data) || strings.TrimSpace(string(data)) == "" {
		c.fail(name, "expected non-empty UTF-8 text")
	}
	return string(data)
}

func policyPath(name string) bool {
	if path.Ext(name) != ".md" || name == "awesome/index.md" {
		return false
	}
	if base := path.Base(name); base == "AGENTS.md" || base == "SKILL.md" {
		return false
	}
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return false
		}
	}
	switch parts[0] {
	case "modules", "examples":
		return len(parts) >= 3
	case "shared", "awesome":
		return len(parts) >= 2
	default:
		return false
	}
}

// Walk without following links. Validate destinations against the resulting set
// of regular documents, so a link cannot enter an unselected module or symlink.
func (c *checker) collect(snapshot bool) {
	err := fs.WalkDir(os.DirFS(c.root), ".", func(name string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !snapshot {
			top, _, _ := strings.Cut(name, "/")
			switch top {
			case ".", "modules", "awesome", "shared", "examples":
			default:
				if entry.IsDir() {
					return fs.SkipDir
				}
				return nil
			}
			if name == "awesome/index.md" || name == "examples/pattern-anti-pattern.md" {
				return nil
			}
			// Other example folders can hold adoption guides and downstream project fixtures.
			if strings.HasPrefix(name, "examples/") {
				fields := strings.Split(name, "/")
				if _, err := os.Lstat(filepath.Join(c.root, "modules", fields[1], "doc.md")); err != nil {
					if entry.IsDir() {
						return fs.SkipDir
					}
					return nil
				}
			}
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() || info.Mode().Perm()&0o111 != 0 {
			c.fail(name, "expected a regular, non-executable file")
			return nil
		}
		if snapshot && name == "LICENSE" {
			return nil
		}
		if !policyPath(name) {
			c.fail(name, "outside the consumer documentation layout")
			return nil
		}
		c.docs[name] = c.read(name)
		return nil
	})
	if err != nil {
		c.fail(c.root, "%v", err)
	}
}

func proseLines(body string, visit func(int, string)) {
	var fence string
	for i, line := range strings.Split(body, "\n") {
		trim := strings.TrimSpace(line)
		if fence != "" {
			if strings.HasPrefix(trim, fence) && strings.Trim(trim, fence[:1]) == "" {
				fence = ""
			}
			continue
		}
		if strings.HasPrefix(trim, "```") || strings.HasPrefix(trim, "~~~") {
			fence = trim[:len(trim)-len(strings.TrimLeft(trim, trim[:1]))]
			continue
		}
		visit(i+1, line)
	}
}

func destination(groups []string) string {
	if groups[1] != "" {
		return groups[1]
	}
	return groups[2]
}

func (c *checker) checkLink(name string, line int, dest string) {
	u, err := url.Parse(dest)
	if err != nil {
		c.fail(name, "line %d: invalid URL %q", line, dest)
		return
	}
	if (u.Scheme == "https" || u.Scheme == "http") && u.Host != "" {
		return // External references are never fetched by this check.
	}
	if u.IsAbs() || u.Host != "" || strings.HasPrefix(u.Path, "/") || strings.Contains(u.Path, "\\") {
		c.fail(name, "line %d: expected a relative policy link or HTTP(S) reference: %q", line, dest)
		return
	}
	if u.Path == "" {
		return // An anchor in the current document.
	}
	resolved := path.Join(path.Dir(name), u.Path)
	if _, ok := c.docs[resolved]; !ok {
		c.fail(name, "line %d: missing consumer document %q (resolved to %s)", line, dest, resolved)
	}
}

func (c *checker) checkLinks() {
	for name, body := range c.docs {
		if strings.HasPrefix(name, "examples/") {
			module := strings.Split(name, "/")[1]
			if _, ok := c.docs["modules/"+module+"/doc.md"]; !ok {
				c.fail(name, "example requires its %s module", module)
			}
		}
		proseLines(body, func(n int, line string) {
			for _, match := range inlineLink.FindAllStringSubmatch(line, -1) {
				c.checkLink(name, n, destination(match))
			}
			if match := definition.FindStringSubmatch(line); match != nil {
				c.checkLink(name, n, destination(match))
			}
		})
	}
}

func (c *checker) checkCatalogs() {
	modules := make(map[string]int)
	proseLines(c.read("README.md"), func(_ int, line string) {
		if row := moduleRow.FindStringSubmatch(strings.TrimSpace(line)); row != nil {
			modules[row[2]]++
			if row[2] != "modules/"+row[1]+"/doc.md" {
				c.fail("README.md", "module key %q does not match %q", row[1], row[2])
			}
		}
	})
	libraries := make(map[string]int)
	for _, line := range strings.Split(c.read("awesome/index.md"), "\n") {
		cols := strings.Split(strings.TrimSpace(line), "|")
		if len(cols) < 4 || cols[0] != "" {
			continue
		}
		target := strings.TrimSpace(cols[2])
		if match := inlineLink.FindStringSubmatch(target); match != nil {
			target = destination(match)
		}
		if strings.HasSuffix(target, ".md") {
			libraries[path.Join("awesome", target)]++
		}
	}
	for name, body := range c.docs {
		if ok, _ := path.Match("modules/*/doc.md", name); ok && modules[name] == 0 {
			c.fail("README.md", "module missing from catalog: %s", name)
		}
		if ok, _ := path.Match("awesome/*.md", name); ok {
			if libraries[name] == 0 {
				c.fail("awesome/index.md", "library list missing from catalog: %s", name)
			}
			c.checkStatuses(name, body)
		}
	}
	for catalog, counts := range map[string]map[string]int{"README.md": modules, "awesome/index.md": libraries} {
		if len(counts) == 0 {
			c.fail(catalog, "empty catalog")
		}
		for name, count := range counts {
			if _, ok := c.docs[name]; !ok || count != 1 {
				c.fail(catalog, "expected one entry pointing to an existing document: %s", name)
			}
		}
	}
}

func (c *checker) checkStatuses(name, body string) {
	rows := 0
	proseLines(body, func(n int, line string) {
		cols := strings.Split(strings.TrimSpace(line), "|")
		if len(cols) < 4 || cols[0] != "" {
			return
		}
		status := strings.TrimSpace(cols[2])
		if status == "status" || strings.Trim(status, ":- ") == "" && strings.Contains(status, "-") {
			return
		}
		if status != "required" && status != "preferred" && status != "banned" {
			c.fail(name, "line %d: invalid library status %q", n, status)
		}
		rows++
	})
	if rows == 0 {
		c.fail(name, "expected library choices with required, preferred or banned status")
	}
}

func (c *checker) checkSkill() {
	entries, err := filepath.Glob(filepath.Join(c.root, "skills", "*", "SKILL.md"))
	if err != nil || len(entries) != 1 || entries[0] != filepath.Join(c.root, skillPath) {
		c.fail("skills", "expected one apply-agent-docs skill entry point")
		return
	}
	body := strings.ReplaceAll(c.read(skillPath), "\r\n", "\n")
	front, rest, ok := strings.Cut(strings.TrimPrefix(body, "---\n"), "\n---\n")
	if !strings.HasPrefix(body, "---\n") || !ok || strings.TrimSpace(rest) == "" {
		c.fail(skillPath, "expected YAML frontmatter and a skill body")
		return
	}
	fields := map[string]string{}
	for _, line := range strings.Split(front, "\n") {
		key, val, _ := strings.Cut(line, ":")
		fields[key] = strings.Trim(strings.TrimSpace(val), "'\"")
	}
	if fields["name"] != "apply-agent-docs" || fields["description"] == "" {
		c.fail(skillPath, "expected a matching name and non-empty description")
	}
}

func validate(root string, snapshot bool) []string {
	c := &checker{root: root, docs: map[string]string{}}
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() {
		return []string{root + ": expected a directory"}
	}
	c.collect(snapshot)
	if snapshot {
		if info, err := os.Lstat(filepath.Join(root, "LICENSE")); err != nil || !info.Mode().IsRegular() {
			c.fail("LICENSE", "include the upstream license")
		}
	}
	for _, name := range []string{"modules/common/doc.md", "modules/taskfile/doc.md"} {
		if _, ok := c.docs[name]; !ok {
			c.fail(name, "baseline module required")
		}
	}
	c.checkLinks()
	if !snapshot {
		c.checkCatalogs()
		c.checkSkill()
	}
	slices.Sort(c.failures)
	return c.failures
}

func main() {
	snapshot := flag.String("snapshot", "", "validate a downstream policy directory and its local links")
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}
	root := "."
	if *snapshot != "" {
		root = *snapshot
	}
	if failures := validate(root, *snapshot != ""); len(failures) > 0 {
		fmt.Fprintln(os.Stderr, strings.Join(failures, "\n"))
		os.Exit(1)
	}
	fmt.Println("Validation passed.")
}
