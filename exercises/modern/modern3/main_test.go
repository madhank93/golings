// modern3
// Go 1.23: for-range can iterate over an iterator function (iter.Seq).

package main_test

import (
	"iter"
	"testing"
)

func countUp(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
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
