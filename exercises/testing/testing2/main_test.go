// testing2
// Make the test pass!
//
// abs is correct. The test fails every time because it reports a failure
// without checking anything first. Make it fail only when abs is wrong.

package main_test

import "testing"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := abs(-3)
	if got != 3 {
		t.Errorf("abs(-3) = %d, want 3", got)
	}
}
