// arrays1
// Arrays are zero-indexed — fix the out-of-range access.
//
// The first element is at index 0 and the last of a length-3 array is at index 2.

package main_test

import "testing"

func first(colors [3]string) string { return colors[0] }
func last(colors [3]string) string  { return colors[2] }

func TestColors(t *testing.T) {
	colors := [3]string{"red", "green", "blue"}
	if got := first(colors); got != "red" {
		t.Errorf(`first: want "red", got %q`, got)
	}
	if got := last(colors); got != "blue" {
		t.Errorf(`last: want "blue", got %q`, got)
	}
}
