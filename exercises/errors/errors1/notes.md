## errors1 — errors are values you return

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Why it works**

- Go has no exceptions: a function that can fail returns an `error` as its last
  result. `errors.New` builds one; `nil` means success, and the caller checks on
  the next line.

**Under the hood**

- `error` is an interface with a single method, `Error() string`. `errors.New`
  returns a pointer to a tiny struct wrapping the message, so two `errors.New`
  calls with identical text are **different** values — comparison is by identity,
  which is exactly what makes sentinel errors work.

**Common mistake**

- Returning a real value alongside a non-nil error and expecting callers to use
  it. The convention is the opposite: when `err != nil`, every other result is
  meaningless. Return the zero value, as here.

**Key detail:** error strings are lowercase and unpunctuated — `"division by
zero"`, not `"Division by zero."` — because they get wrapped into longer
sentences by callers. `go vet`'s `ST1005` check enforces it.

**See also:** errors2 (wrapping) · errors3 (custom types) · functions4 (the
`(T, error)` signature) · the [chapter](../README.md)

**References**

- Effective Go — Errors: https://go.dev/doc/effective_go#errors
- Go Code Review Comments — error strings: https://go.dev/wiki/CodeReviewComments#error-strings
