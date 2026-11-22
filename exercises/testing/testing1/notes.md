## testing1: the test signature

```go
func TestDouble(t *testing.T) {
    if got := double(4); got != 8 {
        t.Errorf("double(4) = %d, want 8", got)
    }
}
```

**Why it works**

- `go test` builds a list of functions named `TestXxx` that take a single
  `*testing.T`. `testing.T` without the `*` names the struct itself, which the
  runner cannot pass, so it refuses to build the package.

**Common mistake**

- Fixing the name instead of the type. `testDouble(t *testing.T)` compiles and
  `go test` reports `ok`, because it ran zero tests. `go test -v` lists each
  test it ran, which makes a missing one easy to spot.

**Key detail:** `*testing.T` is shared state for one test: its failed flag, its
log, its cleanup functions. That is why every helper you write takes the
pointer, never a copy.

**See also:** testing2 · testadv1 (subtests) · the [chapter](../README.md)

**References**

- pkg.go.dev: testing: https://pkg.go.dev/testing
