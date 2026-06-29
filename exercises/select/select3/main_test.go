// select3
// A `default` case makes a select non-blocking.

package main_test

import "testing"

func tryReceive(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
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
		t.Errorf("want ok=false on empty channel")
	}
}
