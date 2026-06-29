// select1
// select waits on multiple channel operations and proceeds with whichever
// one is ready.

package main_test

import "testing"

func firstReady(a, b <-chan int) int {
	select {
	case v := <-a:
		return v
	case v := <-b:
		return v
	}
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
