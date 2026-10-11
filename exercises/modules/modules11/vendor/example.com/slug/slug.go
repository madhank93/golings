// Package slug turns titles into URL-safe identifiers.
package slug

import "strings"

// Make lower-cases title and joins its words with hyphens.
func Make(title string) string {
	return strings.Join(strings.Fields(strings.ToLower(title)), "-")
}
