// context2
// context.WithTimeout cancels itself after a duration.

package main_test

import (
	"context"
	"testing"
	"time"
)

func doWork(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestWorkTimesOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := doWork(ctx, time.Second)
	if err == nil {
		t.Fatal("expected a timeout error")
	}
	if err != context.DeadlineExceeded {
		t.Errorf("want context.DeadlineExceeded, got %v", err)
	}
}

func TestWorkCompletes(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := doWork(ctx, 5*time.Millisecond); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
