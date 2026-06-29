// methods1
// A value receiver works on a copy and only reads; a pointer receiver can
// mutate the original. Implement both.

package main_test

import "testing"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
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
