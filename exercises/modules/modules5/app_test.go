package app_test

import (
	"testing"

	"example.com/app"
)

func TestWeeklyAverage(t *testing.T) {
	days := [7]float64{1, 2, 3, 4, 5, 6, 7}
	if got := app.WeeklyAverage(days); got != 4 {
		t.Errorf("WeeklyAverage = %v, want 4", got)
	}
}
