// stdlib4 — time
// Go formats and parses time using a REFERENCE date, not strftime codes:
//   Mon Jan 2 15:04:05 MST 2006   (i.e. 01/02 03:04:05PM '06 -0700)
// Parse a date string using the layout "2006-01-02".

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// parseDate parses s (formatted as YYYY-MM-DD) into a time.Time.
func parseDate(s string) (time.Time, error) {
	// FIXME: return time.Parse("2006-01-02", s).
	return time.Time{}, nil
}

func TestParseDate(t *testing.T) {
	d, err := parseDate("2026-06-22")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.Year() != 2026 || d.Month() != time.June || d.Day() != 22 {
		t.Errorf("got %v, want 2026-06-22", d.Format("2006-01-02"))
	}
}
