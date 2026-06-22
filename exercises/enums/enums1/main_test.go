// enums1
// Go has no enum keyword; the idiom is a named integer type plus a const block
// using iota, which yields 0, 1, 2, ... for successive constants.

// I AM NOT DONE
package main_test

import "testing"

type Weekday int

// FIXME: declare the rest of the weekdays so Sunday=0 ... Saturday=6.
// After the first `= iota`, each subsequent line repeats it automatically.
const (
	Sunday Weekday = iota
)

func TestWeekdayValues(t *testing.T) {
	if Sunday != 0 {
		t.Errorf("Sunday = %d, want 0", Sunday)
	}
	if Monday != 1 {
		t.Errorf("Monday = %d, want 1", Monday)
	}
	if Saturday != 6 {
		t.Errorf("Saturday = %d, want 6", Saturday)
	}
}
