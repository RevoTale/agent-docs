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
	routerDocPath   = "doc.md"
	rootPolicyPath  = "AGENTS.md"
	moduleDocGlob   = "modules/*/doc.md"
	moduleLegacyGlob = "modules/*/AGENTS.md"

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

func isCodeFence(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "```")
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
