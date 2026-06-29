// defer2
// defer is the idiomatic way to guarantee cleanup runs on every return path.

package main_test

import "testing"

type Resource struct{ closed bool }

func (r *Resource) Close() { r.closed = true }

func process(r *Resource, early bool) {
	defer r.Close()

	if early {
		return
	}
}

func TestDeferCleanup(t *testing.T) {
	r1 := &Resource{}
	process(r1, true)
	if !r1.closed {
		t.Error("resource not closed on the early-return path")
	}

	r2 := &Resource{}
	process(r2, false)
	if !r2.closed {
		t.Error("resource not closed on the normal-return path")
	}
}
