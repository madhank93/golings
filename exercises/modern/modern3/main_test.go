// modern3
// Go 1.23: for-range can iterate over an iterator function. An iter.Seq[V] is
// func(yield func(V) bool) — call yield for each value; stop if it returns
// false. Implement the iterator's body.

// I AM NOT DONE
package main_test

import (
	"iter"
	"testing"
)

// countUp returns an iterator that yields 1, 2, ..., n.
func countUp(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// FIXME: loop 1..n and yield each value; stop early if yield returns false.
	}
}

func TestCountUp(t *testing.T) {
	var got []int
	for v := range countUp(3) {
		got = append(got, v)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("want [1 2 3], got %v", got)
	}
}
