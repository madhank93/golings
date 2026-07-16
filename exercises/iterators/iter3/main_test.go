// iter3
// slices.Collect(seq) drains an iter.Seq[V] into a []V. It does NOT accept an
// iter.Seq2[K, V]; to collect the values of a pair-iterator you first project it
// down to a single-value iterator. valuesOf is that adapter.

package main_test

import (
	"iter"
	"maps"
	"slices"
	"testing"
)

// valuesOf adapts a key/value iterator into a values-only iterator so the
// result can be passed to slices.Collect.
func valuesOf[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		// FIXME: range seq's (key, value) pairs and yield ONLY the value of
		// each, so this returns an iter.Seq[V]. Right now it yields nothing.
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

func TestValuesOf(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := slices.Collect(valuesOf(maps.All(m)))
	slices.Sort(got)
	if want := []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
