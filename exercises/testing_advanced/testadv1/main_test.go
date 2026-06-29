// testadv1 — table-driven tests with t.Run
// Table-driven tests pair a slice of cases with t.Run subtests, so each case
// reports independently. Implement fizzbuzz to make the table pass.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// fizzbuzz returns "Fizz" for multiples of 3, "Buzz" for multiples of 5,
// "FizzBuzz" for multiples of 15, otherwise the number as a string.
func fizzbuzz(n int) string {
	// FIXME: apply the FizzBuzz rules; strconv.Itoa renders the plain number.
	return ""
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
