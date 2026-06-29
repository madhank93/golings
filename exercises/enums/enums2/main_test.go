// enums2
// Pair iota constants with a String() method to get readable names (and free
// pretty-printing via fmt, since this satisfies fmt.Stringer).

// I AM NOT DONE
package main_test

import "testing"

type Color int

const (
	Red Color = iota
	Green
	Blue
)

// String returns the name of the color.
func (c Color) String() string {
	// FIXME: switch on the color and return its name.
	return ""
}

func TestColorString(t *testing.T) {
	cases := map[Color]string{Red: "Red", Green: "Green", Blue: "Blue"}
	for c, want := range cases {
		if got := c.String(); got != want {
			t.Errorf("Color(%d).String() = %q, want %q", int(c), got, want)
		}
	}
}
