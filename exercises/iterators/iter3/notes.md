## iter3 — adapt Seq2 to Seq

```go
func valuesOf[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
    return func(yield func(V) bool) {
        for _, v := range seq { // drop the key, keep the value
            if !yield(v) {
                return
            }
        }
    }
}
slices.Collect(valuesOf(maps.All(m))) // []V of the map's values
```

**Why it works**

- `slices.Collect` drains an `iter.Seq[V]` into a `[]V`, but won't accept a
  pair-iterator (`Seq2`). `valuesOf` **projects** a `Seq2[K,V]` down to a
  `Seq[V]` by ranging the pairs and yielding only the value.

**Key detail:** this is the adapter pattern between the two iterator shapes — the
standard library also ships `maps.Values`/`maps.Keys` for maps specifically.
Writing your own `valuesOf` shows how `Seq` and `Seq2` interconvert so any
pair-source can feed a value-only sink like `Collect`.

**References**

- pkg.go.dev — slices.Collect: https://pkg.go.dev/slices#Collect
- pkg.go.dev — maps.All: https://pkg.go.dev/maps#All
