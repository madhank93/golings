// Package ledger is the bank's public API for balances.
package ledger

import "example.com/bank/ledger/internal/money"

// Format renders a balance in cents, overdrafts included.
func Format(cents int) string {
	if cents < 0 {
		return "-" + money.Dollars(-cents)
	}
	return money.Dollars(cents)
}
