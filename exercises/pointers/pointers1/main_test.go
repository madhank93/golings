// pointers1
// A pointer holds the address of a value. &x takes the address; *p reads or
// writes the value at that address. Modify the caller's int through a pointer.

package main_test

import "testing"

// double multiplies the value that p points to by 2, in place.
func double(p *int) {
	*p = *p * 2
}

func TestDouble(t *testing.T) {
	n := 21
	double(&n)
	if n != 42 {
		t.Errorf("want 42, got %d", n)
	}
}
