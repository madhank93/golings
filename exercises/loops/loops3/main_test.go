// loops3
// Make the test pass!
//
// sumUntilNegative adds the numbers in order, skipping zeros, and stops at
// the first negative number. Use continue and break.

// I AM NOT DONE
package main_test

import "testing"

func sumUntilNegative(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		// FIXME: skip zeros with continue, and stop at a negative with break.
		total += nums[i]
	}
	return total
}

func TestSumUntilNegative(t *testing.T) {
	if got := sumUntilNegative([]int{3, 0, 4, -1, 100}); got != 7 {
		t.Errorf("sumUntilNegative([3 0 4 -1 100]) = %d, want 7", got)
	}
	if got := sumUntilNegative([]int{-5, 1}); got != 0 {
		t.Errorf("sumUntilNegative([-5 1]) = %d, want 0", got)
	}
}
