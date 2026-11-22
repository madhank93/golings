// loops1
// Make the test pass!
//
// sumTo adds up every number from 1 to n. Its loop stops one number too
// early.

// I AM NOT DONE
package main_test

import "testing"

func sumTo(n int) int {
	total := 0
	for i := 1; i < n; i++ {
		total += i
	}
	return total
}

func TestSumTo(t *testing.T) {
	if got := sumTo(4); got != 10 {
		t.Errorf("sumTo(4) = %d, want 10 (1+2+3+4)", got)
	}
	if got := sumTo(1); got != 1 {
		t.Errorf("sumTo(1) = %d, want 1", got)
	}
}
