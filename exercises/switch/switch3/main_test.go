// switch3
// Return a weekday name for a number using a switch.

package main_test

import "testing"

func weekDay(day int) string {
	switch day {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func TestWeekDay(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{0, "Sunday"}, {1, "Monday"}, {2, "Tuesday"}, {3, "Wednesday"},
		{4, "Thursday"}, {5, "Friday"}, {6, "Saturday"},
	}
	for _, tc := range tests {
		if day := weekDay(tc.input); day != tc.want {
			t.Errorf("expected %s but got %s", tc.want, day)
		}
	}
}
