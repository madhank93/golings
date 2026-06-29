// files2 — reading line by line with bufio.Scanner
// bufio.Scanner reads any io.Reader one line at a time.

package main_test

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	n := 0
	for scanner.Scan() {
		n++
	}
	return n
}

func TestCountLines(t *testing.T) {
	r := strings.NewReader("alpha\nbeta\ngamma\n")
	if got := countLines(r); got != 3 {
		t.Errorf("want 3, got %d", got)
	}
}
