// iter4
// A `range` loop drives a push iterator from start to finish — you cannot pause
// one, walk another, and resume. iter.Pull flips a push iterator (iter.Seq)
// into a next/stop pair you call yourself, which is what lockstep work over two
// sequences needs.

package main_test

import (
	"iter"
	"slices"
	"testing"
)

// zip pairs a[i] with b[i], stopping as soon as either sequence runs out.
func zip(a, b iter.Seq[string]) []string {
	var out []string
	next, stop := iter.Pull(b)
	defer stop()
	for x := range a {
		y, ok := next()
		if !ok {
			break
		}
		out = append(out, x+y)
	}
	return out
}

func TestZip(t *testing.T) {
	a := slices.Values([]string{"a", "b", "c"})
	b := slices.Values([]string{"1", "2"})

	got := zip(a, b)
	want := []string{"a1", "b2"}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
