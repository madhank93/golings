// Package pad aligns text in fixed-width columns.
package pad

import "strings"

// Right pads s with spaces to width.
func Right(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
