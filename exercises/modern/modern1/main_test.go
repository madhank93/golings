// modern1
// Go 1.21 added built-in min and max (any ordered type, two or more args) and
// clear (empties a map or zeroes a slice). No import needed.

// I AM NOT DONE
package main_test

import "testing"

// bounds returns the smallest and largest of three ints.
func bounds(a, b, c int) (lo, hi int) {
	// FIXME: use the built-in min and max.
	return 0, 0
}

// wipe empties the map in place.
func wipe(m map[string]int) {
	// FIXME: use the built-in clear.
}

func TestBounds(t *testing.T) {
	lo, hi := bounds(3, 1, 2)
	if lo != 1 || hi != 3 {
		t.Errorf("want lo=1 hi=3, got lo=%d hi=%d", lo, hi)
	}
}

func TestWipe(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	wipe(m)
	if len(m) != 0 {
		t.Errorf("want empty map, got %v", m)
	}
}
