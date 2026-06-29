// select2
// Pairing a receive with time.After in a select implements a timeout.

package main_test

import (
	"testing"
	"time"
)

func recvOrTimeout(ch <-chan int, d time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(d):
		return 0, false
	}
}

func TestRecvGetsValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 9
	if v, ok := recvOrTimeout(ch, time.Second); !ok || v != 9 {
		t.Errorf("want (9,true), got (%d,%v)", v, ok)
	}
}

func TestRecvTimesOut(t *testing.T) {
	ch := make(chan int)
	if _, ok := recvOrTimeout(ch, 10*time.Millisecond); ok {
		t.Errorf("expected a timeout (ok=false)")
	}
}
