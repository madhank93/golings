// stdlib6 — regexp
// regexp matches and extracts text using regular expressions.

// I AM NOT DONE
package main_test

import "testing"

// isEmail reports whether s looks like a basic email: something@something.tld.
func isEmail(s string) bool {
	// FIXME: compile a pattern and match it (import "regexp"):
	//   re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	//   return re.MatchString(s)
	return false
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
