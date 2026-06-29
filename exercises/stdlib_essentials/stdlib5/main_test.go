// stdlib5 — strconv
// strconv converts between strings and numbers.

package main_test

import (
	"strconv"
	"testing"
)

func parseAndDouble(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n * 2, nil
}

func TestParseAndDouble(t *testing.T) {
	got, err := parseAndDouble("21")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("want 42, got %d", got)
	}
	if _, err := parseAndDouble("nope"); err == nil {
		t.Error("expected an error for non-numeric input")
	}
}
