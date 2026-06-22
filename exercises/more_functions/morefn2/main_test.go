// morefn2 — variadic functions
// A variadic parameter (nums ...int) accepts any number of arguments and is
// seen inside the function as a slice. A slice can be spread with slice... .

// I AM NOT DONE
package main_test

import "testing"

// sum returns the total of all its arguments.
func sum(nums ...int) int {
	total := 0
	// FIXME: range over nums and add each value to total.
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
	if got := sum(nums...); got != 15 { // spread a slice into a variadic call
		t.Errorf("sum(nums...) = %d, want 15", got)
	}
}
