// loops2
// Make the test pass!
//
// Go has no while keyword: a for with only a condition is a while loop.
// halvings counts how many times n can be halved before it reaches 1.

package main_test

import "testing"

func halvings(n int) int {
	steps := 0
	for n > 1 {
		n /= 2
		steps++
	}
	return steps
}

func TestHalvings(t *testing.T) {
	cases := map[int]int{1: 0, 2: 1, 8: 3, 10: 3}
	for n, want := range cases {
		if got := halvings(n); got != want {
			t.Errorf("halvings(%d) = %d, want %d", n, got, want)
		}
	}
}
