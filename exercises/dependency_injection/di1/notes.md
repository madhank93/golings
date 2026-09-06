## di1 — inject an io.Writer

```go
func Greet(w io.Writer, name string) {
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

**Why it works**

- `fmt.Fprintf` writes formatted text to any `io.Writer`. Taking the destination
  as a parameter means production can pass `os.Stdout` while the test passes a
  `*bytes.Buffer` and reads back what was written.

**Under the hood**

- `io.Writer` is one method — `Write([]byte) (int, error)` — and files, buffers,
  network connections, gzip wrappers and `httptest` recorders all satisfy it.
  Depending on the smallest possible interface is what makes the substitution
  free.

**Common mistake**

- Reaching for `fmt.Printf` and then trying to capture stdout in the test by
  swapping `os.Stdout`. It works, it is racy under parallel tests, and it is
  unnecessary — the parameter is one word longer and removes the problem.

**Key detail:** note the `Fprintf` / `Printf` / `Sprintf` family: `F` writes to a
writer, no prefix writes to stdout, `S` returns a string. Choosing `F` is the
whole of dependency injection here.

**See also:** di2 (injecting a clock) · di3 (constructor injection) ·
mock1 (asserting on the double) · the [chapter](../README.md)

**References**

- pkg.go.dev — io.Writer: https://pkg.go.dev/io#Writer ·
  fmt.Fprintf: https://pkg.go.dev/fmt#Fprintf
- Go Code Review Comments — interfaces: https://go.dev/wiki/CodeReviewComments#interfaces
