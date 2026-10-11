// Package report prints a stock summary.
package report

import (
	"strconv"

	"example.com/table"
)

// Line renders one item of the summary.
func Line(item string, count int) string {
	return table.Row(item, strconv.Itoa(count), 8)
}
