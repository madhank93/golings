## strings2 — count runes, not bytes

```go
func charCount(s string) int {
    return utf8.RuneCountInString(s)
}
```

**Why it works**

- `len(s)` counts **bytes**, and a non-ASCII character like `é` or `世` takes more
  than one byte in UTF-8. `utf8.RuneCountInString` counts **runes** (code
  points), so `"héllo"` correctly reports 5, not 6.

**Key detail:** a Go string is a read-only slice of **bytes** holding UTF-8 text.
`len` = byte count; ranging with `for i, r := range s` yields **runes**. Reach
for `unicode/utf8` (or `range`) whenever "length" should mean characters.

**References**

- unicode/utf8 package: https://pkg.go.dev/unicode/utf8
- The Go Blog — Strings, bytes, runes and characters in Go: https://go.dev/blog/strings
