// context2
// context.WithTimeout cancels itself after a duration.
// Make doWork abort early and return ctx.Err() when the context is done.

// I AM NOT DONE
package main_test

import (
	"context"
	"testing"
	"time"
)

// doWork simulates work that takes d, but must stop early and return ctx.Err()
// if the context is cancelled first. Returns nil when the work finishes in time.
func doWork(ctx context.Context, d time.Duration) error {
	// FIXME: replace the blocking sleep with a select that waits on EITHER
	// time.After(d) (work done -> return nil) OR ctx.Done() (return ctx.Err()).
	time.Sleep(d)
	return nil
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
