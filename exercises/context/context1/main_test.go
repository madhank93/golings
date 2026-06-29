// context1
// A context.Context lets a caller signal cancellation to running work.
// Make countUntilCancelled return once the context is cancelled.

// I AM NOT DONE
package main_test

import (
	"context"
	"testing"
	"time"
)

// countUntilCancelled increments until ctx is cancelled, then returns the count.
func countUntilCancelled(ctx context.Context) int {
	count := 0
	for {
		select {
		// FIXME: add a case that returns when the context is done.
		case <-time.After(time.Millisecond):
			count++
		}
	}
}

func TestCountStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	done := make(chan int, 1)
	go func() { done <- countUntilCancelled(ctx) }()

	select {
	case <-done:
		// returned after cancellation — good
	case <-time.After(2 * time.Second):
		t.Fatal("countUntilCancelled never returned after cancel")
	}
}
