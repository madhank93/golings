// unsafe1
// unsafe.Offsetof(x.Field) reports the byte offset of Field within its struct —
// layout introspection with zero runtime cost and no reflection.

// I AM NOT DONE
package main_test

import (
	"testing"
	"unsafe"
)

type Record struct {
	A byte
	B int64
	C int32
}

// offsetOfC must return the byte offset of field C within Record.
func offsetOfC() uintptr {
	// FIXME: return the real offset with unsafe.Offsetof(Record{}.C) instead of
	// this placeholder.
	return 0
}

func TestOffsetOfC(t *testing.T) {
	var r Record
	want := unsafe.Offsetof(r.C)
	if got := offsetOfC(); got != want {
		t.Errorf("want offset %d, got %d", want, got)
	}
}
