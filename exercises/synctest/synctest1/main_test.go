// synctest1
// testing/synctest (Go 1.25) runs test code in a "bubble". synctest.Wait()
// blocks until every goroutine in the bubble is idle — a deterministic
// replacement for time.Sleep-and-hope. But Wait only works inside a bubble;
// call it outside one and it panics. The body must be wrapped in synctest.Test.

// I AM NOT DONE
package main_test

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
)

func TestBubbleWait(t *testing.T) {
	// FIXME: this body calls synctest.Wait(), which only works inside a bubble,
	// so it panics as written. Wrap everything below in:
	//   synctest.Test(t, func(t *testing.T) { ... })
	var count atomic.Int64
	for range 3 {
		go func() { count.Add(1) }()
	}
	synctest.Wait()
	if count.Load() != 3 {
		t.Errorf("want 3, got %d", count.Load())
	}
}
