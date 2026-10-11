package shop_test

import (
	"testing"

	"example.com/shop"
)

func TestCheckout(t *testing.T) {
	if got := shop.Checkout(1200, 300); got != 2000 {
		t.Errorf("Checkout(1200, 300) = %d, want 2000", got)
	}
}
