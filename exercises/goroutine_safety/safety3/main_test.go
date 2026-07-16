// safety3
// Channel ownership: the goroutine that OWNS a channel closes it — exactly once.
// Consumers must never close a channel they only receive from. Several workers
// each closing the same channel panics with "close of closed channel".

// I AM NOT DONE
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

// drain fills a channel, closes it once (as the owner), then fans the values
// out to three worker goroutines that sum them.
func drain(nums []int) int64 {
	ch := make(chan int, len(nums))
	for _, n := range nums {
		ch <- n
	}
	close(ch) // owner closes the channel, once, after the last send

	var total atomic.Int64
	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func() {
			for v := range ch {
				total.Add(int64(v))
			}
			// FIXME: a worker only RECEIVES from ch — it does not own it and
			// must not close it. With three workers each calling close, the
			// second one panics "close of closed channel". Delete this line;
			// the owner already closed ch above.
			close(ch)
		})
	}
	wg.Wait()
	return total.Load()
}

func TestDrain(t *testing.T) {
	if got := drain([]int{1, 2, 3, 4, 5}); got != 15 {
		t.Errorf("want 15, got %d", got)
	}
}
