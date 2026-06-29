// channels1
// A buffered channel — make(chan T, n) — holds up to n values without a
// receiver ready.

package main_test

import "testing"

func collect(vals []int) []int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)

	out := make([]int, 0, len(vals))
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func TestCollect(t *testing.T) {
	got := collect([]int{1, 2, 3})
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("want [1 2 3], got %v", got)
	}
}
