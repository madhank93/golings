// di1
// Inject an io.Writer so the output can be captured and asserted in a test.

// I AM NOT DONE
package main_test

import (
	"bytes"
	"io"
	"testing"
)

// Greet writes "Hello, <name>!" to w. Real code can pass os.Stdout; a test
// passes a *bytes.Buffer and inspects what was written.
func Greet(w io.Writer, name string) {
	// FIXME: write the greeting to w — which fmt function writes formatted text to an io.Writer?
	_ = w
}

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	Greet(&buf, "Go")
	if got, want := buf.String(), "Hello, Go!"; got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
