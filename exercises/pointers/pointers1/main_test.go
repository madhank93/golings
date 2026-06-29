// pointers1
// A pointer holds the address of a value. &x takes the address; *p reads or
// writes the value at that address. Modify the caller's int through a pointer.

// I AM NOT DONE
package main_test

import "testing"

// double multiplies the value that p points to by 2, in place.
func double(p *int) {
	// FIXME: write through the pointer to change the caller's value.
}

func TestDouble(t *testing.T) {
	n := 21
	double(&n)
	if n != 42 {
		t.Errorf("want 42, got %d", n)
	}
}
