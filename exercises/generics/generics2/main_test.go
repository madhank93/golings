// generics2
// Constrain a generic to add numbers of any numeric type.
//
// A type constraint lists the types a parameter may take; a generic function
// then works for every one of them while keeping the same body.

package main_test

import "testing"

type Number interface {
	int | float64
}

func addNumbers[T Number](n1, n2 T) T {
	return n1 + n2
}

func TestAddNumbers(t *testing.T) {
	if got := addNumbers(1, 2); got != 3 {
		t.Errorf("want 3, got %v", got)
	}
	if got := addNumbers(1.5, 2.5); got != 4.0 {
		t.Errorf("want 4, got %v", got)
	}
}
