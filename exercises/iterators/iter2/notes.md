## iter2 — Enumerate with iter.Seq2

```go
func Enumerate(s []int) iter.Seq2[int, int] {
    return func(yield func(int, int) bool) {
        for i, v := range s {
            if !yield(i, v) { // (index, value) — in that order
                return
            }
        }
    }
}
```

**Why it works**

- `iter.Seq2[K, V]` = `func(yield func(K, V) bool)` — an iterator over **pairs**,
  the two-value analog of `iter.Seq`. `Enumerate` yields `(i, v)`, so ranging it
  gives index first, value second — just like `for i, v := range slice`.

**Key detail:** order matters — `yield(i, v)` not `yield(v, i)`. `Seq2` is what
powers `for k, v := range` over maps and custom pair iterators. The consumer
writes `for i, v := range Enumerate(s)`, mirroring the built-in slice range.

**References**

- pkg.go.dev — iter.Seq2: https://pkg.go.dev/iter#Seq2
