// concpat2 — fan-in
// Fan-in merges values from several input channels into one output channel.
// Implement merge so it copies every input into out and closes out when done.

// I AM NOT DONE
package main_test

import (
	"sort"
	"sync"
	"testing"
)

// merge fans several channels into one, closing the result once all are drained.
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// FIXME: for each input channel c, wg.Add(1) and start a goroutine that
	// copies c's values into out then calls wg.Done(). Then start one more
	// goroutine that does wg.Wait() followed by close(out).

	return out
}

// gen returns a buffered channel emitting vals, already closed.
func gen(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestFanIn(t *testing.T) {
	var got []int
	for v := range merge(gen(1, 2, 3), gen(4, 5)) {
		got = append(got, v)
	}
	sort.Ints(got)

	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("want %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}
