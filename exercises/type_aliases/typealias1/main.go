// typealias1
// A defined type like `type Celsius float64` is DISTINCT from float64 — Go will
// not mix the two without an explicit conversion, which stops accidental unit
// mix-ups at compile time.

// I AM NOT DONE
package main

import "fmt"

type Celsius float64

// toFahrenheit takes a plain float64.
func toFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	var boiling Celsius = 100
	// FIXME: toFahrenheit wants a float64, but boiling is a Celsius. Convert it
	// explicitly with float64(boiling) at the call site.
	fmt.Printf("%.0f°F\n", toFahrenheit(boiling))
}
