// sync3
// sync/atomic provides lock-free atomic operations on integers.

package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func countHits(n int) int64 {
	var hits int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&hits, 1)
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&hits)
}

func TestCountHits(t *testing.T) {
	if got := countHits(1000); got != 1000 {
		t.Errorf("want 1000, got %d", got)
	}
}
