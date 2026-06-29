// sync1
// A sync.Mutex guards shared state so concurrent goroutines don't race.
// Run with `go test -race` to see the unsynchronized version flagged.

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
)

// Counter is incremented concurrently and must stay correct.
type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	// FIXME: take the lock before touching n and release it after (defer).
	c.n++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func TestConcurrentInc(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("want 1000, got %d", got)
	}
}
