// Package units converts between measurement systems. Version 2 of the module
// returns the converted value together with its unit symbol.
package units

// CToF converts degrees Celsius to Fahrenheit.
func CToF(c float64) (float64, string) {
	return c*9/5 + 32, "°F"
}
