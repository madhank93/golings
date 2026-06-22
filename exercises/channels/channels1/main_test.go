// channels1
// A buffered channel — make(chan T, n) — holds up to n values without a
// receiver ready. An unbuffered channel blocks the sender until someone
// receives, which would deadlock the single-goroutine loop below.

// I AM NOT DONE
package main_test

import "testing"

// collect sends all vals into a channel, closes it, then drains it back out.
func collect(vals []int) []int {
	// FIXME: make the channel buffered with capacity len(vals) so the sends
	// in the loop don't block without a concurrent receiver:
	//   ch := make(chan int, len(vals))
	ch := make(chan int) // unbuffered -> deadlocks when sending below
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
