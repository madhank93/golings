// Package app reports on a week of scores.
package app

import "example.com/mathx"

// WeeklyAverage averages seven daily scores.
func WeeklyAverage(days [7]float64) float64 {
	return mathx.Mean(days[:]...)
}
