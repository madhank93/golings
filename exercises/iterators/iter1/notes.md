## iter1 — Filter an iter.Seq

```go
func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
    return func(yield func(V) bool) {
        for v := range seq {
            if !keep(v) {
                continue // drop values that fail the predicate
            }
            if !yield(v) {
                return
            }
        }
    }
}
```

**Why it works**

- `Filter` wraps a source sequence and returns a **new** `iter.Seq` that only
  re-yields values where `keep(v)` is true. Ranging the source and guarding the
  `yield` with the predicate is the whole trick.

**Key detail:** iterators **compose** — `Filter` takes a `Seq` and returns a `Seq`, so
you can chain `Map`, `Filter`, `Take`, etc. Nothing runs until something ranges
the final sequence (lazy), and no intermediate slices are allocated. Propagate
`yield`'s `false` return so early `break`s stop the source too.

**References**

- pkg.go.dev — iter: https://pkg.go.dev/iter
- The Go Blog — Range Over Function Types: https://go.dev/blog/range-functions
