// interfaces2
// Implementing fmt.Stringer (String() string) controls how a value prints.
// The `var _ fmt.Stringer = p` line is a compile-time check that Point
// satisfies the interface.

// I AM NOT DONE
package main_test

import (
	"fmt"
	"testing"
)

type Point struct{ X, Y int }

// FIXME: implement String() returning the point as "(X, Y)".

func TestStringer(t *testing.T) {
	var _ fmt.Stringer = Point{} // compile-time check: Point must satisfy Stringer

	p := Point{X: 3, Y: 4}
	if got := fmt.Sprintf("%v", p); got != "(3, 4)" {
		t.Errorf("want (3, 4), got %s", got)
	}
}
