## strings1 — compose helpers from the strings package

```go
func slugify(s string) string {
    return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
```

**Why it works**

- `strings.ToLower` lowercases the text; its result is fed to
  `strings.ReplaceAll`, which swaps every space for a hyphen. Reading inside-out:
  lowercase first, then replace.

**Key detail:** Go strings are **immutable** — every `strings` helper returns a *new*
string rather than editing in place. That's why you chain and return the result
instead of mutating `s`.

**References**

- strings package: https://pkg.go.dev/strings
- Go by Example — String Functions: https://gobyexample.com/string-functions
