// methods1
// Methods attach behavior to a type. A value receiver (r Rectangle) works on a
// copy and can only read; a pointer receiver (r *Rectangle) can mutate the
// original. Implement both.

// I AM NOT DONE
package main_test

import "testing"

type Rectangle struct {
	Width, Height float64
}

// Area returns Width*Height. Value receiver — it only reads.
func (r Rectangle) Area() float64 {
	// FIXME: return r.Width * r.Height
	return 0
}

// Scale multiplies both sides by factor IN PLACE. Pointer receiver — it mutates.
func (r *Rectangle) Scale(factor float64) {
	// FIXME: multiply r.Width and r.Height by factor.
}

func TestArea(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if r.Area() != 12 {
		t.Errorf("want 12, got %v", r.Area())
	}
}

func TestScale(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	if r.Width != 6 || r.Height != 8 {
		t.Errorf("want 6x8, got %vx%v", r.Width, r.Height)
	}
	if r.Area() != 48 {
		t.Errorf("want area 48 after scaling, got %v", r.Area())
	}
}
