// channels2
// Channel direction in a signature documents and enforces intent:
// chan<- T is send-only, <-chan T is receive-only.

// I AM NOT DONE
package main_test

import (
	"sort"
	"testing"
)

// produce sends 0..n-1 on out, then closes it. out is SEND-ONLY.
func produce(out chan<- int, n int) {
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
}

// consume reads everything from in (RECEIVE-ONLY) and returns the values.
func consume(in <-chan int) []int {
	var got []int
	// FIXME: range over in and append each value to got.
	return got
}

func TestProduceConsume(t *testing.T) {
	ch := make(chan int)
	go produce(ch, 5)

	got := consume(ch)
	sort.Ints(got)
	if len(got) != 5 || got[0] != 0 || got[4] != 4 {
		t.Errorf("want 0..4, got %v", got)
	}
}
