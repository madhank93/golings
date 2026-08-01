## stdlib6 — regexp

```go
re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
re.MatchString(s)
```

**Why it works**

- `regexp.MustCompile` builds a reusable pattern; `MatchString` reports whether
  `s` matches. The anchors `^`/`$` force the whole string to match, and
  `[^@\s]+` means "one or more non-`@`, non-space characters".

**Key detail:** use `MustCompile` for **static** patterns (it panics on a bad regex at
startup, surfacing typos immediately); use `regexp.Compile` (returns an error) for
patterns built from user input. Compile **once** and reuse — recompiling inside a
loop is wasteful. Go uses RE2 syntax (linear-time, no catastrophic backtracking).

**References**

- pkg.go.dev — regexp: https://pkg.go.dev/regexp
- RE2 syntax: https://pkg.go.dev/regexp/syntax
