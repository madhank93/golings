package tui

import _ "embed"

// helpText is the `?` pop-up. It lives in a file rather than a string literal
// because it is markdown full of backticks, and a raw Go string cannot contain
// one.
//
//go:embed help.md
var helpText string
