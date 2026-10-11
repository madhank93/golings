// Package table renders rows as aligned columns.
package table

import "example.com/pad"

// Row renders two cells, the first padded to width.
func Row(label, value string, width int) string {
	return pad.Right(label, width) + value
}
