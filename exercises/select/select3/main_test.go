// select3
// A `default` case makes a select non-blocking: it runs when no other
// case is ready instead of waiting.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// tryReceive returns (value, true) if ch has a value ready right now,
// otherwise (0, false) without blocking.
func tryReceive(ch <-chan int) (int, bool) {
	// FIXME: use a select with `case v := <-ch:` and a `default:` case
	// so it returns immediately when nothing is ready.
	v := <-ch
	return v, true
}

func TestTryReceiveReady(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 5
	if v, ok := tryReceive(ch); !ok || v != 5 {
		t.Errorf("want (5,true), got (%d,%v)", v, ok)
	}
}

func TestTryReceiveEmpty(t *testing.T) {
	ch := make(chan int)
	if _, ok := tryReceive(ch); ok {
		t.Errorf("expected (_,false) on an empty channel")
	}
}
