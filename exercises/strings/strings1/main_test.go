// strings1
// The strings package holds the everyday string utilities.

package main_test

import (
	"strings"
	"testing"
)

// slugify lowercases s and replaces every space with a hyphen.
func slugify(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

func TestSlugify(t *testing.T) {
	if got := slugify("Hello World"); got != "hello-world" {
		t.Errorf("want hello-world, got %q", got)
	}
	if got := slugify("Go Is Fun"); got != "go-is-fun" {
		t.Errorf("want go-is-fun, got %q", got)
	}
}
