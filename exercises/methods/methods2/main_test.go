// methods2
// You can define methods on ANY named type you declare, not just structs.

package main_test

import "testing"

type Celsius float64

func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func TestFahrenheit(t *testing.T) {
	if got := Celsius(100).Fahrenheit(); got != 212 {
		t.Errorf("100C: want 212F, got %v", got)
	}
	if got := Celsius(0).Fahrenheit(); got != 32 {
		t.Errorf("0C: want 32F, got %v", got)
	}
}
