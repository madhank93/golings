// modern2
// Go 1.22: a for-range loop can iterate over an integer directly —
// `for i := range n` runs with i = 0, 1, ..., n-1.

// I AM NOT DONE
package main_test

import "testing"

// sumTo returns 0 + 1 + ... + (n-1) using range over an integer.
func sumTo(n int) int {
	total := 0
	// FIXME: for i := range n { total += i }
	return total
}

func TestSumTo(t *testing.T) {
	if got := sumTo(5); got != 10 {
		t.Errorf("want 10, got %d", got)
	}
	if got := sumTo(0); got != 0 {
		t.Errorf("want 0, got %d", got)
	}
}
