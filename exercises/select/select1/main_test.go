// select1
// select waits on multiple channel operations and proceeds with whichever
// one is ready. Return the first value available from either channel.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// firstReady returns the first value available from a or b.
func firstReady(a, b <-chan int) int {
	// FIXME: use a select with a `case v := <-a:` and a `case v := <-b:`,
	// returning v from whichever is ready.
	return 0
}

func TestFirstReadyFromA(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int)
	a <- 42
	if got := firstReady(a, b); got != 42 {
		t.Errorf("want 42, got %d", got)
	}
}

func TestFirstReadyFromB(t *testing.T) {
	a := make(chan int)
	b := make(chan int, 1)
	b <- 7
	if got := firstReady(a, b); got != 7 {
		t.Errorf("want 7, got %d", got)
	}
}
