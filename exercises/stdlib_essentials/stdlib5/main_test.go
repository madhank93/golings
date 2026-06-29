// stdlib5 — strconv
// strconv converts between strings and numbers (and other base types).

// I AM NOT DONE
package main_test

import "testing"

// parseAndDouble parses s as an int and returns it doubled. It returns an
// error when s is not a valid integer.
func parseAndDouble(s string) (int, error) {
	// FIXME: convert the string to an int (strconv), handle the error, then double it.
	return 0, nil
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
