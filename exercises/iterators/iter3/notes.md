## iter3 — project a Seq2 down to a Seq

```go
func valuesOf[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
    return func(yield func(V) bool) {
        for _, v := range seq {
            if !yield(v) {
                return
            }
        }
    }
}

slices.Collect(valuesOf(maps.All(m)))
```

**Why it works**

- `slices.Collect` takes an `iter.Seq[V]`, not an `iter.Seq2[K, V]`. `valuesOf`
  ranges the pairs and re-yields only the value half, producing the single-value
  iterator `Collect` accepts.

**Under the hood**

- Ranging a `Seq2` inside the adapter gives both halves; discarding the key with
  `_` is what performs the projection. The adapter stays lazy — it is still just
  a function waiting to be called.

**Common mistake**

- Reaching for `maps.Values(m)` and being surprised the adapter exists at all.
  It does exist for maps — but the general case (any `Seq2`, including one
  produced by your own code) needs this three-line projection, which is why it is
  worth being able to write.

**Key detail:** the same shape gives you `keysOf` (yield `k`), a swap
(`yield(v, k)`), or a filter over pairs. Adapters between `Seq` and `Seq2` are
the connective tissue of iterator pipelines.

**See also:** iter1 (`Filter`) · iter2 (pair order) · mapspkg1 (`maps.Keys` /
`maps.Values`) · the [chapter](../README.md)

**References**

- pkg.go.dev — slices.Collect: https://pkg.go.dev/slices#Collect
- pkg.go.dev — maps.All: https://pkg.go.dev/maps#All
