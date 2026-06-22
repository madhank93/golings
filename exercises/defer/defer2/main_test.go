// defer2
// defer is the idiomatic way to guarantee cleanup runs on every return path
// (the classic `f, _ := os.Open(...); defer f.Close()` pattern).

// I AM NOT DONE
package main_test

import "testing"

type Resource struct{ closed bool }

func (r *Resource) Close() { r.closed = true }

// process must Close r before returning, whether it returns early or not.
func process(r *Resource, early bool) {
	// FIXME: defer r.Close() at the top so it runs on EVERY return path below.

	if early {
		return
	}
	// ... imagine more work here ...
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
