// iter1
// iter.Seq[V] is func(yield func(V) bool) — the shape of a range-over-func
// iterator. A Filter wraps a source sequence and re-yields only the values that
// match a predicate.

// I AM NOT DONE
package main_test

import (
	"iter"
	"slices"
	"testing"
)

// Filter returns an iterator over the elements of seq for which keep(v) is true.
func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			// FIXME: only forward v when keep(v) is true. As written every
			// element is yielded, so nothing is filtered out. Guard the yield
			// with the predicate.
			if !yield(v) {
				return
			}
		}
	}
}

func TestFilter(t *testing.T) {
	src := slices.Values([]int{1, 2, 3, 4, 5, 6})
	even := Filter(src, func(n int) bool { return n%2 == 0 })
	got := slices.Collect(even)
	want := []int{2, 4, 6}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
