// stdlib4 — time
// Go formats and parses time using a REFERENCE date: 2006-01-02 15:04:05.

package main_test

import (
	"testing"
	"time"
)

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
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
