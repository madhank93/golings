// iter2
// iter.Seq2[K, V] is func(yield func(K, V) bool) — an iterator over pairs, like
// ranging a map or a slice's index/value. Enumerate pairs each element with its
// index, so the FIRST value yielded is the index and the SECOND is the element.

// I AM NOT DONE
package main_test

import (
	"iter"
	"slices"
	"testing"
)

// Enumerate yields (index, value) for each element of s.
func Enumerate(s []int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, v := range s {
			// FIXME: the pair is yielded in the wrong order. It should be
			// (index, value) — yield(i, v) — not (value, index).
			if !yield(v, i) {
				return
			}
		}
	}
}

func TestEnumerate(t *testing.T) {
	var idxs, vals []int
	for i, v := range Enumerate([]int{10, 20, 30}) {
		idxs = append(idxs, i)
		vals = append(vals, v)
	}
	if !slices.Equal(idxs, []int{0, 1, 2}) || !slices.Equal(vals, []int{10, 20, 30}) {
		t.Errorf("want idx [0 1 2] val [10 20 30], got idx %v val %v", idxs, vals)
	}
}
