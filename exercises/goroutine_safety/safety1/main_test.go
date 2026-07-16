// safety1
// A value shared across goroutines needs synchronization. sync.RWMutex allows
// many concurrent readers (RLock) but an exclusive writer (Lock) — a good fit
// for read-heavy state. With no lock at all, concurrent access is a data race.

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
)

// Counter is incremented and read from many goroutines at once.
type Counter struct {
	// FIXME: add a sync.RWMutex field and guard Inc with Lock/Unlock and Value
	// with RLock/RUnlock. Without it, `go test -race` reports a data race and
	// the final count is wrong.
	n int
}

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Value() int {
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
