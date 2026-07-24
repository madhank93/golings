// arrays1
// Arrays are zero-indexed — fix the out-of-range access.
//
// The first element is at index 0 and the last of a length-3 array is at index 2.

// I AM NOT DONE
package main_test

import "testing"

// FIXME: return the first and last colors using the right indexes.
func first(colors [3]string) string { return colors[] }
func last(colors [3]string) string  { return colors[] }

func TestColors(t *testing.T) {
	colors := [3]string{"red", "green", "blue"}
	if got := first(colors); got != "red" {
		t.Errorf(`first: want "red", got %q`, got)
	}
	if got := last(colors); got != "blue" {
		t.Errorf(`last: want "blue", got %q`, got)
	}
}
