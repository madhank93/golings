// generics4
// Write a generic Reduce that folds any slice into a single accumulated value.

package main_test

import "testing"

func Reduce[A, B any](items []A, init B, f func(B, A) B) B {
	acc := init
	for _, item := range items {
		acc = f(acc, item)
	}
	return acc
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
