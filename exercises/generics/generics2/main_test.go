// generics2
// Constrain a generic to add numbers of any numeric type.
//
// A type constraint lists the types a parameter may take; a generic function
// then works for every one of them while keeping the same body.

// I AM NOT DONE
package main_test

import "testing"

// FIXME: make Number allow ints AND floats, then give addNumbers a type
// parameter T constrained by Number and a return type.
type Number interface {
	int
}

func addNumbers(n1, n2 T) {
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
