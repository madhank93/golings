// modules2
// An import path is the module path from go.mod followed by the package's
// directory inside the module. Checkout needs the cart package in ./cart —
// fix the import so the toolchain can find it.

package shop

import "example.com/shop/cart"

// Shipping is the flat fee added to every order, in cents.
const Shipping = 500

// Checkout returns the order total in cents, shipping included.
func Checkout(prices ...int) int {
	return cart.Total(prices...) + Shipping
}
