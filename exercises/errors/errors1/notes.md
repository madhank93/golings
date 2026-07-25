## errors1 — errors are values

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

**Why it works**

- Go has **no exceptions** for ordinary failures. A function returns an `error`
  as its last value; `nil` means success, non-`nil` means failure.

**Nuance:** the caller must **check** `err` before trusting the result — the
`result, err := f()` then `if err != nil` pattern is everywhere in Go. Return a
zero result alongside a non-nil error so callers never use a half-built value.

**References**

- Go by Example — Errors: https://gobyexample.com/errors
- The Go Blog — Errors are values: https://go.dev/blog/errors-are-values
