// stdlib3 — slices
// The slices package (Go 1.21+) provides generic helpers over slices.
// Sort the numbers in DESCENDING order.

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

// sortDesc returns a new slice with nums sorted from largest to smallest.
func sortDesc(nums []int) []int {
	out := slices.Clone(nums)
	// FIXME: sort out descending. slices.Sort sorts ascending — sort then
	// slices.Reverse(out), or use slices.SortFunc with a comparison.
	return out
}

func TestSortDesc(t *testing.T) {
	got := sortDesc([]int{3, 1, 2, 5, 4})
	want := []int{5, 4, 3, 2, 1}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
