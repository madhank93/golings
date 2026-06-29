// stdlib3 — slices
// The slices package (Go 1.21+) provides generic helpers over slices.

package main_test

import (
	"slices"
	"testing"
)

func sortDesc(nums []int) []int {
	out := slices.Clone(nums)
	slices.Sort(out)
	slices.Reverse(out)
	return out
}

func TestSortDesc(t *testing.T) {
	got := sortDesc([]int{3, 1, 2, 5, 4})
	want := []int{5, 4, 3, 2, 1}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
