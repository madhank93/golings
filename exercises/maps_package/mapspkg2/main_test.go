// mapspkg2
// maps.DeleteFunc(m, pred) (Go 1.21) removes every entry for which pred(k, v) is
// true, in a single pass — no collect-the-keys-then-delete two-step needed.

package main_test

import (
	"maps"
	"testing"
)

// dropExpired removes every entry whose value (a TTL) is <= 0.
func dropExpired(m map[string]int) {
	// FIXME: this collects the keys to remove but never deletes them. Replace
	// the whole loop with one call (add import "maps"):
	//   maps.DeleteFunc(m, func(k string, v int) bool { return v <= 0 })
	maps.DeleteFunc(m, func(k string, v int) bool { return v <= 0 })
}

func TestDropExpired(t *testing.T) {
	m := map[string]int{"a": 5, "b": 0, "c": 3, "d": -1}
	dropExpired(m)
	want := map[string]int{"a": 5, "c": 3}
	if len(m) != len(want) {
		t.Fatalf("want %v, got %v", want, m)
	}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("want %v, got %v", want, m)
		}
	}
}
