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
	// FIXME: sort ascending then reverse — or sort with a custom comparison.
	return out
}

func TestSortDesc(t *testing.T) {
	got := sortDesc([]int{3, 1, 2, 5, 4})
	want := []int{5, 4, 3, 2, 1}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
