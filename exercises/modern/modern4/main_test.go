// modern4
// Go 1.25: WaitGroup.Go(f) launches f in its own goroutine and takes care of
// Add(1) and Done() for you. It replaces the classic — and easy to get wrong —
// wg.Add(1); go func() { defer wg.Done(); ... }() dance.

package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

// sumParallel adds each number from its own goroutine, then waits for them all.
func sumParallel(nums []int) int64 {
	var wg sync.WaitGroup
	var total atomic.Int64
	for _, n := range nums {
		// FIXME: launch the work with wg.Go(func() { ... }) instead of a bare
		// `go`. As written nothing calls wg.Add, so wg.Wait() returns before the
		// goroutines run and total is read too early.
		wg.Go(func() {
			total.Add(int64(n))
		})
	}
	wg.Wait()
	return total.Load()
}

func TestSumParallel(t *testing.T) {
	if got := sumParallel([]int{1, 2, 3, 4, 5}); got != 15 {
		t.Errorf("want 15, got %d", got)
	}
}
