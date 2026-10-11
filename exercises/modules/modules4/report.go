// modules4
// A package inside an internal/ directory can only be imported by code rooted
// at internal's parent. money lives in ledger/internal/, so only ledger may
// use it. Format the balance through ledger's public API instead.

package bank

import "example.com/bank/ledger"

// Statement returns the balance line printed at the bottom of a statement.
func Statement(cents int) string {
	return "Balance: " + ledger.Format(cents)
}
