// modules6
// From v2 on, a module's major version is part of its path: go.mod already
// requires example.com/units/v2. The import path has to say so too — the
// package name inside stays plain units.

// I AM NOT DONE
package weather

import (
	"fmt"

	"example.com/units"
)

// Forecast renders a temperature for an American audience.
func Forecast(celsius float64) string {
	f, unit := units.CToF(celsius)
	return fmt.Sprintf("%.0f%s", f, unit)
}
