## strings1 — the strings package

```go
func slugify(s string) string {
    return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
```

**Why it works**

- `ToLower` lowercases every letter and `ReplaceAll` swaps each space for a
  hyphen. Both **return new strings** — a Go string is immutable, so no function
  in the package edits its input.

**Under the hood**

- Each call allocates: a string header is a pointer plus a length, and the bytes
  behind it can never be written to, so a "modification" is always a fresh
  buffer. Two calls means two allocations, which is fine here and matters in a
  loop.

**Common mistake**

- Building strings with `+=` inside a loop. Every `+=` copies the whole
  accumulated string, so n appends copy O(n²) bytes. Use `strings.Builder`,
  which grows one buffer and hands it over with `String()`.

**Key detail:** for case-insensitive *comparison* use `strings.EqualFold(a, b)`
rather than `ToLower(a) == ToLower(b)` — no allocations and proper Unicode
folding.

**See also:** strings2 (bytes vs runes) · primitive_types2 (format verbs) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — strings: https://pkg.go.dev/strings ·
  strings.Builder: https://pkg.go.dev/strings#Builder
- Go by Example — String Functions: https://gobyexample.com/string-functions
