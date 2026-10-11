// modules3
// Only names that start with an upper-case letter are visible outside their
// package. The tests are in a separate package, temp_test, so they see exactly
// what any other importer would. Export what they use.

// I AM NOT DONE
package temp

// Reading is one thermometer sample.
type Reading struct {
	celsius float64
}

// fahrenheit converts a reading for display.
func fahrenheit(r Reading) float64 {
	return r.celsius*9/5 + 32
}
