// Package oldfmt is the formatter the report used before table replaced it.
package oldfmt

// Row joins cells with a pipe.
func Row(cells ...string) string {
	out := ""
	for i, c := range cells {
		if i > 0 {
			out += "|"
		}
		out += c
	}
	return out
}
