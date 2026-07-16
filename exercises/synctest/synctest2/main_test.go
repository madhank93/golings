// synctest2
// Inside a synctest bubble time is virtual: when EVERY goroutine (including the
// test's own) is blocked, the fake clock jumps to the next timer. So a one-hour
// time.Sleep finishes instantly — but only once the test goroutine actually
// blocks waiting for it. Read the clock without waiting and no time passes.

// I AM NOT DONE
package main_test

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestVirtualClock(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		done := make(chan struct{})
		go func() {
			time.Sleep(time.Hour)
			close(done)
		}()

		// FIXME: nothing here blocks, so the bubble never advances its clock and
		// elapsed is 0. Receive from the channel — add `<-done` — so both
		// goroutines block, virtual time jumps an hour, and the sleeper finishes.
		if elapsed := time.Since(start); elapsed != time.Hour {
			t.Errorf("want exactly 1h of virtual time, got %v", elapsed)
		}
	})
}
