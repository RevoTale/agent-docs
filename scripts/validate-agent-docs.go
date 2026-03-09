package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	routerDocPath    = "doc.md"
	rootPolicyPath   = "AGENTS.md"
	moduleDocGlob    = "modules/*/doc.md"
	moduleLegacyGlob = "modules/*/AGENTS.md"
	awesomeIndexPath = "awesome/index.md"
	awesomeDocGlob   = "awesome/*.md"
	skillsDirPath    = "skills"
	skillDocGlob     = "skills/*/SKILL.md"

	sectionStrictRules      = "Strict rules"
	sectionWorkingAgreement = "Working Agreements"

	scannerInitialCapacity = 64 * 1024
	scannerMaxCapacity     = 1024 * 1024
)

var (
	headingPattern      = regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
	bulletPattern       = regexp.MustCompile(`^\s*-\s+`)
	normativeBullet     = regexp.MustCompile(`^\s*-\s+(MUST|SHOULD|MAY)(\s|$)`)
	registryRowPattern  = regexp.MustCompile(`^\|\s*[^|]+\|\s*[^|]+\|\s*modules/[^|]+\|`)
	awesomeRowPattern   = regexp.MustCompile(`^\|\s*[^|]+\|\s*(\[[^]]+\]\([^)]*\.md\)|\.?/?.*\.md)\s*\|`)
	awesomeStatusRow    = regexp.MustCompile(`^\|\s*[^|]+\|\s*(required|preferred|banned)\s*\|`)
	skillNamePattern    = regexp.MustCompile(`^name:\s*(.+?)\s*$`)
	skillDescPattern    = regexp.MustCompile(`^description:\s*(.+?)\s*$`)
	waForbiddenPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)task\s+(gen(:check|:code-diff)?|fix|validate|test)(\s|$)`),
		regexp.MustCompile(`(?i)bun\s+(run|install)(\s|$)`),
		regexp.MustCompile(`(?i)go\s+test(\s|$)`),
		regexp.MustCompile(`(?i)golangci-lint`),
		regexp.MustCompile(`(?i)gofmt`),
		regexp.MustCompile(`(?i)tsc(\s|$)`),
		regexp.MustCompile(`(?i)must\s+pass`),
		regexp.MustCompile(`(?i)before\s+merge`),
	}
	expectedTopLevelSections = []string{
		"Overview",
		"Project structure",
		sectionStrictRules,
		"Working Agreements",
	}
)

type validator struct {
	failures []string
}

type skillDoc struct {
	name        string
	description string
	body        string
}

func (v *validator) failf(format string, args ...any) {
	v.failures = append(v.failures, fmt.Sprintf(format, args...))
}

func (v *validator) report() error {
	if len(v.failures) == 0 {
		return nil
	}

	for _, failure := range v.failures {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", failure)
	}
	return fmt.Errorf("validation failed with %d error(s)", len(v.failures))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0, 128)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, scannerInitialCapacity), scannerMaxCapacity)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func parseHeading(line string) (level int, title string, ok bool) {
	matches := headingPattern.FindStringSubmatch(line)
	if matches == nil {
		return 0, "", false
	}

	return len(matches[1]), strings.TrimSpace(matches[2]), true
}

func extractSection(lines []string, title string) ([]string, bool) {
	inSection := false
	sectionLines := make([]string, 0, 32)

	for _, line := range lines {
		_, headingTitle, isHeading := parseHeading(line)
		if isHeading {
			if inSection {
				break
			}
			if headingTitle == title {
				inSection = true
			}
			continue
		}
		if inSection {
			sectionLines = append(sectionLines, line)
		}
	}

	return sectionLines, inSection
}

func parseRegistryPaths(path string) ([]string, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, 16)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !registryRowPattern.MatchString(trimmed) {
			continue
		}

		parts := strings.Split(trimmed, "|")
		if len(parts) < 5 {
			continue
		}

		modulePath := strings.TrimSpace(parts[3])
		if modulePath == "" {
			continue
		}
		paths = append(paths, filepath.Clean(modulePath))
	}

	return paths, nil
}

func parseAwesomeRegistryPaths(path string) ([]string, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, 8)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !awesomeRowPattern.MatchString(trimmed) {
			continue
		}

		parts := strings.Split(trimmed, "|")
		if len(parts) < 4 {
			continue
		}

		awesomePath := strings.TrimSpace(parts[2])
		if awesomePath == "" {
			continue
		}

		awesomePath = strings.TrimPrefix(awesomePath, "[")
		if closeIdx := strings.Index(awesomePath, "]("); closeIdx >= 0 {
			target := awesomePath[closeIdx+2:]
			awesomePath = strings.TrimSuffix(target, ")")
		}

		clean := filepath.Clean(filepath.Join("awesome", strings.TrimPrefix(awesomePath, "./")))
		paths = append(paths, clean)
	}

	return paths, nil
}

func isCodeFence(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "```")
}

func parseSkillDoc(path string) (skillDoc, error) {
	lines, err := readLines(path)
	if err != nil {
		return skillDoc{}, err
	}

	if len(lines) < 3 || strings.TrimSpace(lines[0]) != "---" {
		return skillDoc{}, fmt.Errorf("missing YAML frontmatter")
	}

	frontmatterEnd := -1
	for idx := 1; idx < len(lines); idx++ {
		if strings.TrimSpace(lines[idx]) == "---" {
			frontmatterEnd = idx
			break
		}
	}
	if frontmatterEnd == -1 {
		return skillDoc{}, fmt.Errorf("frontmatter closing delimiter not found")
	}

	doc := skillDoc{
		body: strings.Join(lines[frontmatterEnd+1:], "\n"),
	}

	for _, line := range lines[1:frontmatterEnd] {
		trimmed := strings.TrimSpace(line)
		if matches := skillNamePattern.FindStringSubmatch(trimmed); matches != nil {
			doc.name = strings.Trim(strings.TrimSpace(matches[1]), `"'`)
			continue
		}
		if matches := skillDescPattern.FindStringSubmatch(trimmed); matches != nil {
			doc.description = strings.Trim(strings.TrimSpace(matches[1]), `"'`)
		}
	}

	return doc, nil
}

func checkRequiredSubstrings(v *validator, path string, required ...string) {
	lines, err := readLines(path)
	if err != nil {
		v.failf("%s cannot be read: %v", path, err)
		return
	}

	fullText := strings.Join(lines, "\n")
	for _, needle := range required {
		if !strings.Contains(fullText, needle) {
			v.failf("%s must mention %q", path, needle)
		}
	}
}

func checkNormativeBullets(v *validator, filePath string, sectionTitle string) {
	lines, err := readLines(filePath)
	if err != nil {
		v.failf("%s cannot be read: %v", filePath, err)
		return
	}

	sectionLines, found := extractSection(lines, sectionTitle)
	if !found {
		return
	}

	inCodeBlock := false
	for _, line := range sectionLines {
		if isCodeFence(line) {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock || !bulletPattern.MatchString(line) {
			continue
		}

		if !normativeBullet.MatchString(line) {
			v.failf("%s has non-normative bullet in %s: %s", filePath, sectionTitle, strings.TrimSpace(line))
		}
	}
}

func checkWorkingAgreementSemantics(v *validator, filePath string) {
	lines, err := readLines(filePath)
	if err != nil {
		v.failf("%s cannot be read: %v", filePath, err)
		return
	}

	sectionLines, found := extractSection(lines, sectionWorkingAgreement)
	if !found {
		return
	}

	inCodeBlock := false
	for _, line := range sectionLines {
		if isCodeFence(line) {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			continue
		}

		for _, pattern := range waForbiddenPatterns {
			if pattern.MatchString(line) {
				v.failf("%s has technical command/gate patterns in Working Agreements: %s", filePath, strings.TrimSpace(line))
				return
			}
		}
	}
}

func checkModuleSectionLayout(v *validator, modulePath string) {
	lines, err := readLines(modulePath)
	if err != nil {
		v.failf("%s cannot be read: %v", modulePath, err)
		return
	}

	topLevelTitles := make([]string, 0, len(expectedTopLevelSections))
	for _, line := range lines {
		level, title, isHeading := parseHeading(line)
		if isHeading && level == 1 {
			topLevelTitles = append(topLevelTitles, title)
		}
	}

	if len(topLevelTitles) != len(expectedTopLevelSections) {
		v.failf("%s must contain exactly %d top-level sections; found %d", modulePath, len(expectedTopLevelSections), len(topLevelTitles))
		return
	}

	for idx, expected := range expectedTopLevelSections {
		if topLevelTitles[idx] != expected {
			v.failf("%s has invalid top-level section order or names", modulePath)
			return
		}
	}
}

func checkAwesomeFile(v *validator, awesomePath string) {
	lines, err := readLines(awesomePath)
	if err != nil {
		v.failf("%s cannot be read: %v", awesomePath, err)
		return
	}

	statusRows := 0
	for _, line := range lines {
		if awesomeStatusRow.MatchString(strings.TrimSpace(line)) {
			statusRows++
		}
	}

	if statusRows == 0 {
		v.failf("%s must include at least one library row with status required|preferred|banned", awesomePath)
	}
}

func checkSkillsLayout(v *validator) {
	entries, err := os.ReadDir(skillsDirPath)
	if err != nil {
		v.failf("%s cannot be read: %v", skillsDirPath, err)
		return
	}

	if len(entries) == 0 {
		v.failf("%s must contain at least one skill directory", skillsDirPath)
		return
	}

	for _, entry := range entries {
		entryPath := filepath.Join(skillsDirPath, entry.Name())
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if !entry.IsDir() {
			v.failf("%s must contain only skill directories; found file: %s", skillsDirPath, entryPath)
			continue
		}

		skillPath := filepath.Join(entryPath, "SKILL.md")
		if _, statErr := os.Stat(skillPath); statErr != nil {
			v.failf("Skill directory is missing SKILL.md: %s", skillPath)
		}
	}
}

func checkSkillFile(v *validator, skillPath string) {
	doc, err := parseSkillDoc(skillPath)
	if err != nil {
		v.failf("%s is invalid: %v", skillPath, err)
		return
	}

	folderName := filepath.Base(filepath.Dir(skillPath))
	if doc.name == "" {
		v.failf("%s must define frontmatter field: name", skillPath)
	}
	if doc.description == "" {
		v.failf("%s must define frontmatter field: description", skillPath)
	}
	if doc.name != "" && doc.name != folderName {
		v.failf("%s frontmatter name must match its folder name (%s)", skillPath, folderName)
	}

	fullText := strings.Join([]string{doc.name, doc.description, doc.body}, "\n")
	lowerText := strings.ToLower(fullText)

	if strings.Contains(fullText, "AGENTS.router.md") {
		v.failf("%s must reference doc.md instead of deprecated AGENTS.router.md", skillPath)
	}

	switch folderName {
	case "init-project-from-agent-docs":
		if !strings.Contains(lowerText, "interview") {
			v.failf("%s must require an architecture interview", skillPath)
		}
		if !strings.Contains(lowerText, "nested") {
			v.failf("%s must describe nested AGENTS.md creation", skillPath)
		}
		if !strings.Contains(fullText, "Accept") {
			v.failf("%s must require explicit Accept before writing", skillPath)
		}
	case "refresh-project-agents-from-agent-docs":
		if !strings.Contains(lowerText, "signal") {
			v.failf("%s must describe repository-signal based selection", skillPath)
		}
		if !strings.Contains(lowerText, "nested") {
			v.failf("%s must describe nested AGENTS.md refresh", skillPath)
		}
		if !strings.Contains(lowerText, "frontend/") && !strings.Contains(lowerText, "apps/*") {
			v.failf("%s must describe nested app or service boundaries", skillPath)
		}
		if !strings.Contains(fullText, "Accept") {
			v.failf("%s must require explicit Accept before writing", skillPath)
		}
	case "refactor-project-to-agent-docs":
		if !strings.Contains(lowerText, "interview") {
			v.failf("%s must require an architecture interview", skillPath)
		}
		if !strings.Contains(lowerText, "subtree") && !strings.Contains(lowerText, "subproject") {
			v.failf("%s must describe per-subtree or per-subproject refactoring", skillPath)
		}
		if !strings.Contains(lowerText, "plan") {
			v.failf("%s must require a refactor plan before edits", skillPath)
		}
		if !strings.Contains(fullText, "Accept") {
			v.failf("%s must require explicit Accept before refactoring", skillPath)
		}
	}
}

func main() {
	v := &validator{}

	legacyFiles, err := filepath.Glob(moduleLegacyGlob)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: legacy glob failure: %v\n", err)
		os.Exit(1)
	}
	sort.Strings(legacyFiles)
	for _, legacy := range legacyFiles {
		v.failf("Legacy module file still exists: %s", legacy)
	}

	registryPaths, err := parseRegistryPaths(routerDocPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s cannot be parsed: %v\n", routerDocPath, err)
		os.Exit(1)
	}
	if len(registryPaths) == 0 {
		v.failf("No module registry paths found in %s canonical table", routerDocPath)
	}

	if _, err := os.Stat(awesomeIndexPath); err != nil {
		v.failf("Awesome registry entrypoint not found: %s", awesomeIndexPath)
	}

	awesomeRegistryPaths, err := parseAwesomeRegistryPaths(awesomeIndexPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s cannot be parsed: %v\n", awesomeIndexPath, err)
		os.Exit(1)
	}
	if len(awesomeRegistryPaths) == 0 {
		v.failf("No awesome files found in %s table", awesomeIndexPath)
	}

	registrySet := make(map[string]struct{}, len(registryPaths))
	for _, registryPath := range registryPaths {
		if _, exists := registrySet[registryPath]; exists {
			v.failf("Duplicate module registry path: %s", registryPath)
			continue
		}
		registrySet[registryPath] = struct{}{}

		if _, statErr := os.Stat(registryPath); statErr != nil {
			v.failf("Registry path not found: %s", registryPath)
		}
	}

	moduleDocs, err := filepath.Glob(moduleDocGlob)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: module glob failure: %v\n", err)
		os.Exit(1)
	}
	sort.Strings(moduleDocs)

	for _, moduleDoc := range moduleDocs {
		if _, exists := registrySet[moduleDoc]; !exists {
			v.failf("Module doc is missing from %s registry: %s", routerDocPath, moduleDoc)
		}
		checkModuleSectionLayout(v, moduleDoc)
	}

	awesomeDocs, err := filepath.Glob(awesomeDocGlob)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: awesome glob failure: %v\n", err)
		os.Exit(1)
	}
	sort.Strings(awesomeDocs)

	awesomeSet := make(map[string]struct{}, len(awesomeRegistryPaths))
	for _, awesomePath := range awesomeRegistryPaths {
		if _, exists := awesomeSet[awesomePath]; exists {
			v.failf("Duplicate awesome registry path: %s", awesomePath)
			continue
		}
		awesomeSet[awesomePath] = struct{}{}
		if _, statErr := os.Stat(awesomePath); statErr != nil {
			v.failf("Awesome registry path not found: %s", awesomePath)
			continue
		}
		checkAwesomeFile(v, awesomePath)
	}

	for _, awesomeDoc := range awesomeDocs {
		if awesomeDoc == awesomeIndexPath {
			continue
		}
		if _, exists := awesomeSet[awesomeDoc]; !exists {
			v.failf("Awesome file is missing from %s registry: %s", awesomeIndexPath, awesomeDoc)
		}
	}

	checkSkillsLayout(v)

	skillDocs, err := filepath.Glob(skillDocGlob)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: skill glob failure: %v\n", err)
		os.Exit(1)
	}
	sort.Strings(skillDocs)

	if len(skillDocs) == 0 {
		v.failf("No skill files found in %s", skillDocGlob)
	}

	for _, skillPath := range skillDocs {
		checkSkillFile(v, skillPath)
	}

	checkRequiredSubstrings(v, routerDocPath, "nested `AGENTS.md`", "nearest `AGENTS.md`", "root `AGENTS.md`")
	checkRequiredSubstrings(v, rootPolicyPath, "nested `AGENTS.md`", "subprojects", "multiple app or service boundaries")
	checkRequiredSubstrings(v, "README.md", "nested `AGENTS.md`", "frontend/", "backend/")

	filesToCheck := make([]string, 0, len(moduleDocs)+2)
	filesToCheck = append(filesToCheck, rootPolicyPath, routerDocPath)
	filesToCheck = append(filesToCheck, moduleDocs...)

	for _, filePath := range filesToCheck {
		checkNormativeBullets(v, filePath, sectionStrictRules)
		checkNormativeBullets(v, filePath, sectionWorkingAgreement)
		checkWorkingAgreementSemantics(v, filePath)
	}

	if err := v.report(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Validation passed.")
}
