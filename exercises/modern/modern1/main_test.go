// modern1
// Go 1.21 added built-in min, max, and clear. No import needed.

package main_test

import "testing"

func bounds(a, b, c int) (lo, hi int) {
	return min(a, b, c), max(a, b, c)
}

func wipe(m map[string]int) {
	clear(m)
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
