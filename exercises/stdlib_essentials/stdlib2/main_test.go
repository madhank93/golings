// stdlib2 — io.Reader / io.Writer
// io.Copy streams bytes from any Reader to any Writer.

package main_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func readAll(r io.Reader) (string, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		return "", err
	}
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
