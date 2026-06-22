// select2
// Pairing a receive with time.After in a select implements a timeout.

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// recvOrTimeout waits up to d for a value on ch.
// Returns (value, true) if one arrives, or (0, false) on timeout.
func recvOrTimeout(ch <-chan int, d time.Duration) (int, bool) {
	// FIXME: select between `case v := <-ch:` (return v, true) and
	// `case <-time.After(d):` (return 0, false).
	v := <-ch
	return v, true
}

func TestRecvGetsValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 9
	if v, ok := recvOrTimeout(ch, time.Second); !ok || v != 9 {
		t.Errorf("want (9,true), got (%d,%v)", v, ok)
	}
}

func TestRecvTimesOut(t *testing.T) {
	ch := make(chan int) // nothing is ever sent
	if _, ok := recvOrTimeout(ch, 10*time.Millisecond); ok {
		t.Errorf("expected a timeout (ok=false)")
	}
}
