// unsafe2
// unsafe.String(ptr, len) builds a string that SHARES the backing array of a
// []byte — a zero-copy conversion for hot paths, unlike string(b) which
// allocates and copies. The bytes must not be mutated afterwards.

package main_test

import (
	"testing"
	"unsafe"
)

// bytesToString converts b to a string WITHOUT copying — the result must share
// b's backing array.
func bytesToString(b []byte) string {
	// FIXME: string(b) allocates and copies. Use
	// unsafe.String(unsafe.SliceData(b), len(b)) so the string shares b's
	// backing array (a true zero-copy conversion).
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func TestBytesToStringZeroCopy(t *testing.T) {
	b := []byte("payload")
	s := bytesToString(b)
	if s != "payload" {
		t.Fatalf("want payload, got %q", s)
	}
	if unsafe.StringData(s) != unsafe.SliceData(b) {
		t.Error("expected a zero-copy conversion sharing b's backing array")
	}
}
