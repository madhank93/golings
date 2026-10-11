// Package cart adds up what is in a shopping cart.
package cart

// Total returns the sum of prices, in cents.
func Total(prices ...int) int {
	sum := 0
	for _, p := range prices {
		sum += p
	}
	return sum
}
