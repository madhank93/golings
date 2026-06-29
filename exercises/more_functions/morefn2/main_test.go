// morefn2 — variadic functions
// A variadic parameter (nums ...int) accepts any number of arguments.

package main_test

import "testing"

// sum returns the total of all its arguments.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func TestSum(t *testing.T) {
	if got := sum(1, 2, 3); got != 6 {
		t.Errorf("sum(1,2,3) = %d, want 6", got)
	}
	if got := sum(); got != 0 {
		t.Errorf("sum() = %d, want 0", got)
	}
	nums := []int{4, 5, 6}
	if got := sum(nums...); got != 15 {
		t.Errorf("sum(nums...) = %d, want 15", got)
	}
}
