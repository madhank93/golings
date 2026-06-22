// sync2
// sync.Once runs a function exactly once, even when called from many goroutines.

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
)

// Config loads its data lazily, but the load must happen only once.
type Config struct {
	once  sync.Once
	loads int
}

func (c *Config) Load() {
	// FIXME: wrap the body in c.once.Do(func() { ... }) so that loads is
	// incremented exactly once across all calls.
	c.loads++
}

func TestLoadOnce(t *testing.T) {
	c := &Config{}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Load()
		}()
	}
	wg.Wait()
	if c.loads != 1 {
		t.Errorf("want loads==1, got %d", c.loads)
	}
}
