# Iterators (iter.Seq)

Range-over-function iterators (Go 1.23) let you write lazy sequences that plug
into `for range` without channels. `iter.Seq[V]` is `func(yield func(V) bool)`
and `iter.Seq2[K, V]` yields pairs. Combined with `slices` and `maps` helpers
they compose into clean data pipelines.

- **iter1** — `iter.Seq[V]`: a lazy `Filter` built from `yield`
- **iter2** — `iter.Seq2[K, V]`: yielding `(index, value)` pairs
- **iter3** — projecting a `Seq2` down to a `Seq` for `slices.Collect`

## Resources

- [iter package](https://pkg.go.dev/iter)
- [Go blog: Range Over Function Types](https://go.dev/blog/range-functions)
- [slices.Collect / slices.Values](https://pkg.go.dev/slices#Collect)
