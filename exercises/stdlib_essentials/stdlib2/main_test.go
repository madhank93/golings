// stdlib2 — io.Reader / io.Writer
// These two interfaces are the backbone of Go I/O. io.Copy streams bytes from
// any Reader to any Writer. Drain the reader into the buffer.

// I AM NOT DONE
package main_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// readAll copies everything from r into a buffer and returns it as a string.
func readAll(r io.Reader) (string, error) {
	var buf bytes.Buffer
	// FIXME: stream the reader into the buffer in one call (io.Copy).
	return buf.String(), nil
}

func TestReadAll(t *testing.T) {
	got, err := readAll(strings.NewReader("hello world"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello world" {
		t.Errorf("want %q, got %q", "hello world", got)
	}
}
