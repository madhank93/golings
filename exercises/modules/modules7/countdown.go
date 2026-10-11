// Package countdown launches rockets.
package countdown

import "strconv"

// From counts down from n to 1 and then shouts.
func From(n int) []string {
	out := make([]string, 0, n+1)
	for i := range n {
		out = append(out, strconv.Itoa(n-i))
	}
	return append(out, "liftoff!")
}
