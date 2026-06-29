// morefn1 — recursion
// A function may call itself. Every recursion needs a base case to stop.
// Implement factorial recursively.

// I AM NOT DONE
package main_test

import "testing"

// factorial returns n! (with 0! == 1).
func factorial(n int) int {
	// FIXME: base case n<=1 returns 1; otherwise recurse on n-1.
	return 0
}

func TestFactorial(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 5: 120, 6: 720}
	for in, want := range cases {
		if got := factorial(in); got != want {
			t.Errorf("factorial(%d) = %d, want %d", in, got, want)
		}
	}
}
