// testadv1 — table-driven tests with t.Run
// Table-driven tests pair a slice of cases with t.Run subtests.

package main_test

import (
	"strconv"
	"testing"
)

func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"multiple of three", 3, "Fizz"},
		{"multiple of five", 5, "Buzz"},
		{"multiple of fifteen", 15, "FizzBuzz"},
		{"plain number", 7, "7"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := fizzbuzz(c.in); got != c.want {
				t.Errorf("fizzbuzz(%d) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
