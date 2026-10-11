// modules3
// Only names that start with an upper-case letter are visible outside their
// package. The tests are in a separate package, temp_test, so they see exactly
// what any other importer would. Export what they use.

package temp

// Reading is one thermometer sample.
type Reading struct {
	Celsius float64
}

// Fahrenheit converts a reading for display.
func Fahrenheit(r Reading) float64 {
	return r.Celsius*9/5 + 32
}
