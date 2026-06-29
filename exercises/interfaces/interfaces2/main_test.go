// interfaces2
// Implementing fmt.Stringer (String() string) controls how a value prints.

package main_test

import (
	"fmt"
	"testing"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func TestStringer(t *testing.T) {
	var _ fmt.Stringer = Point{}

	p := Point{X: 3, Y: 4}
	if got := fmt.Sprintf("%v", p); got != "(3, 4)" {
		t.Errorf("want (3, 4), got %s", got)
	}
}
