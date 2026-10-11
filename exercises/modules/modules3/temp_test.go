package temp_test

import (
	"testing"

	"example.com/temp"
)

func TestFahrenheit(t *testing.T) {
	r := temp.Reading{Celsius: 100}
	if got := temp.Fahrenheit(r); got != 212 {
		t.Errorf("Fahrenheit(100°C) = %v, want 212", got)
	}
}
