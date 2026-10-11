// Package app welcomes visitors.
package app

import "example.com/greet"

// Welcome is the line shown at the front desk.
func Welcome(name string) string {
	return greet.Formal(name)
}
