package exercises

import (
	"os"
	"regexp"
	"strings"
)

// genericComment matches the boilerplate header lines that aren't descriptive.
var genericComment = regexp.MustCompile(`(?i)^make( me compile| the tests? pass|.*pass)`)

// Description returns a one-line summary of what the exercise is about, taken
// from the first meaningful comment in the exercise file (the line after the
// "// <name>" header). Returns "" when there's nothing useful.
func (e Exercise) Description() string {
	data, err := os.ReadFile(e.Path)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "//") {
			if trimmed == "" {
				continue // allow blank lines above/within the header
			}
			break // reached code; stop
		}
		c := strings.TrimSpace(strings.TrimPrefix(trimmed, "//"))
		if c == "" || strings.EqualFold(c, e.Name) || notDoneRegex.MatchString(line) || genericComment.MatchString(c) {
			continue
		}
		return c
	}
	return ""
}
