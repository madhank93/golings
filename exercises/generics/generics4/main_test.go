// generics4
// Write a generic Reduce that folds any slice into a single accumulated value.

// I AM NOT DONE
package main_test

import "testing"

// Reduce folds items into an accumulator: starting from init, it applies f to
// the running accumulator and each element in turn, and returns the result.
func Reduce[A, B any](items []A, init B, f func(B, A) B) B {
	// FIXME: range over items, updating acc := f(acc, item), then return acc.
	return init
}

func TestReduceSum(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
	if sum != 10 {
		t.Errorf("want 10, got %d", sum)
	}
}

func TestReduceConcat(t *testing.T) {
	got := Reduce([]string{"a", "b", "c"}, "", func(acc, s string) string { return acc + s })
	if got != "abc" {
		t.Errorf("want \"abc\", got %q", got)
	}
}
