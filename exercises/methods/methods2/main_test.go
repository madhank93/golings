// methods2
// You can define methods on ANY named type you declare, not just structs.

// I AM NOT DONE
package main_test

import "testing"

// Celsius is a named type whose underlying type is float64.
type Celsius float64

// Fahrenheit converts the temperature to degrees Fahrenheit: c*9/5 + 32.
func (c Celsius) Fahrenheit() float64 {
	// FIXME: convert the receiver to float64 and apply the C→F formula.
	return 0
}

func TestFahrenheit(t *testing.T) {
	if got := Celsius(100).Fahrenheit(); got != 212 {
		t.Errorf("100C: want 212F, got %v", got)
	}
	if got := Celsius(0).Fahrenheit(); got != 32 {
		t.Errorf("0C: want 32F, got %v", got)
	}
}
