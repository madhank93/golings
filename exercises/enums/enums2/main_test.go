// enums2
// Pair iota constants with a String() method to get readable names.

package main_test

import "testing"

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Unknown"
	}
}

func TestColorString(t *testing.T) {
	cases := map[Color]string{Red: "Red", Green: "Green", Blue: "Blue"}
	for c, want := range cases {
		if got := c.String(); got != want {
			t.Errorf("Color(%d).String() = %q, want %q", int(c), got, want)
		}
	}
}
