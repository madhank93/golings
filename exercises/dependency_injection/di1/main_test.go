// di1
// Inject an io.Writer so the output can be captured and asserted in a test.

package main_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s!", name)
}

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	Greet(&buf, "Go")
	if got, want := buf.String(), "Hello, Go!"; got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
