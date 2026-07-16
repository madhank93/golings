// safety1
// A value shared across goroutines needs synchronization. sync.RWMutex allows
// many concurrent readers (RLock) but an exclusive writer (Lock) — a good fit
// for read-heavy state. With no lock at all, concurrent access is a data race.

package main_test

import (
	"sync"
	"testing"
)

// Counter is incremented and read from many goroutines at once.
type Counter struct {
	mu sync.RWMutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.n
}

func TestCounter(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	for range 100 {
		wg.Go(func() {
			c.Inc()
			_ = c.Value()
		})
	}
	wg.Wait()
	if c.Value() != 100 {
		t.Errorf("want 100, got %d", c.Value())
	}
}
