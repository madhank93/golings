// Package money formats amounts. It is internal to ledger: only code under
// ledger/ may import it.
package money

import "fmt"

// Dollars formats a non-negative amount of cents as dollars.
func Dollars(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}
