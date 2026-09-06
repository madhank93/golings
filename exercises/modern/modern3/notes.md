## modern3 — write an iterator function

```go
func countUp(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := 1; i <= n; i++ {
            if !yield(i) {
                return // the consumer stopped
            }
        }
    }
}
```

**Why it works**

- Go 1.23 lets `for range` drive a function of the shape
  `func(yield func(V) bool)` — that is exactly `iter.Seq[int]`. Calling `yield`
  hands one value to the loop body.

**Under the hood**

- The compiler rewrites `for v := range countUp(3)` into a call to your function,
  passing the **loop body itself** as `yield`. Control ping-pongs: the iterator
  calls the body, the body returns `true` to continue. No goroutine, no channel,
  no allocation per element.

**Common mistake**

- Ignoring `yield`'s return value. `false` means the consumer left — a `break`, a
  `return`, an error path — and an iterator that keeps yielding panics with
  `range function continued iteration after loop body exit`.

**Key detail:** this replaces the three old ways of exposing a sequence: building
the whole slice, handing out a channel (a goroutine plus a leak if the consumer
leaves early), or a bespoke `Next()`/`Err()` interface.

**See also:** iter1 (composing iterators) · iter4 (`iter.Pull`) ·
anonymous_functions3 (closures) · the [chapter](../README.md)

**References**

- Go blog — Range Over Function Types: https://go.dev/blog/range-functions
- pkg.go.dev — iter: https://pkg.go.dev/iter
