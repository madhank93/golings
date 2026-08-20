## iter1 — a lazy Filter

```go
func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
    return func(yield func(V) bool) {
        for v := range seq {
            if !keep(v) {
                continue // drop it
            }
            if !yield(v) {
                return // consumer stopped
            }
        }
    }
}
```

**Why it works**

- `Filter` takes an iterator and returns one, ranging the source and re-yielding
  only the values that pass. The broken version yields every element, so nothing
  is filtered.

**Under the hood**

- Nothing is materialised. `Filter` computes a value only when the consumer asks
  for the next one, so `Filter(hugeSeq, …)` costs nothing until it is ranged —
  and a `break` in the consumer stops the source mid-flight.

**Common mistake**

- Guarding only the predicate and forgetting `yield`'s result. Both checks are
  required: `keep` decides what to forward, `yield`'s `false` decides whether to
  keep producing at all.

**Key detail:** `slices.Values(s)` turns a slice into a `Seq` and
`slices.Collect(seq)` drains one back into a slice — the two ends of the
pipeline this exercise sits in the middle of.

**See also:** iter2 (`Seq2`) · iter3 (projection) · modern3 (writing the
protocol) · generics4 (the eager equivalent) · the [chapter](../README.md)

**References**

- pkg.go.dev — iter.Seq: https://pkg.go.dev/iter#Seq
- Go blog — Range Over Function Types: https://go.dev/blog/range-functions
