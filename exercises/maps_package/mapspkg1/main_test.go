// mapspkg1
// The maps package (Go 1.21+) has helpers for whole-map operations. maps.Clone
// returns an independent shallow copy — unlike assigning a map value, where both
// names refer to the SAME underlying map and a write through one is seen by the
// other.

package main_test

import (
	"maps"
	"testing"
)

// withDefault returns a copy of settings with "theme" added, WITHOUT modifying
// the caller's map.
func withDefault(settings map[string]string) map[string]string {
	// FIXME: `cp := settings` does not copy a map — both names point at the same
	// map, so cp["theme"]=... mutates the caller's original. Replace it with
	// cp := maps.Clone(settings) (add import "maps").
	cp := maps.Clone(settings)
	cp["theme"] = "dark"
	return cp
}

func TestWithDefault(t *testing.T) {
	original := map[string]string{"lang": "en"}
	withDefault(original)
	if _, leaked := original["theme"]; leaked {
		t.Error("withDefault must not modify the caller's map; use maps.Clone")
	}
}
