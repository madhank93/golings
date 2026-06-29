// context1
// A context.Context lets a caller signal cancellation to running work.

package main_test

import (
	"context"
	"testing"
	"time"
)

func countUntilCancelled(ctx context.Context) int {
	count := 0
	for {
		select {
		case <-ctx.Done():
			return count
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
	case <-time.After(2 * time.Second):
		t.Fatal("countUntilCancelled never returned after cancel")
	}
}
