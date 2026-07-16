// typealias2
// Methods can be attached to a defined type. Give Celsius a String() method and
// the fmt package will use it automatically (the fmt.Stringer interface).

// I AM NOT DONE
package main_test

import (
	"fmt"
	"testing"
)

type Celsius float64

// FIXME: implement String() on Celsius so it formats like "25°C". Use
// float64(c) inside — calling fmt.Sprint(c) would recurse forever.
//   func (c Celsius) String() string { return fmt.Sprintf("%g°C", float64(c)) }

func TestCelsiusString(t *testing.T) {
	got := fmt.Sprint(Celsius(25))
	if got != "25°C" {
		t.Errorf("want 25°C, got %q", got)
	}
}
