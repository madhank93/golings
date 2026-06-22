// files2 — reading line by line with bufio.Scanner
// bufio.Scanner reads any io.Reader one line at a time, without loading the
// whole input into memory — the usual way to process files or stdin.

// I AM NOT DONE
package main_test

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

// countLines returns the number of lines read from r.
func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	n := 0
	// FIXME: advance the scanner line by line: for scanner.Scan() { n++ }
	return n
}

func TestCountLines(t *testing.T) {
	r := strings.NewReader("alpha\nbeta\ngamma\n")
	if got := countLines(r); got != 3 {
		t.Errorf("want 3, got %d", got)
	}
}
