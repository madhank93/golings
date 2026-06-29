// stdlib6 — regexp
// regexp matches and extracts text using regular expressions.

package main_test

import (
	"regexp"
	"testing"
)

func isEmail(s string) bool {
	re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	return re.MatchString(s)
}

func TestIsEmail(t *testing.T) {
	if !isEmail("a@b.com") {
		t.Error("a@b.com should match")
	}
	if isEmail("nope") {
		t.Error("nope should not match")
	}
	if isEmail("a@b") {
		t.Error("a@b should not match (no top-level domain)")
	}
}
