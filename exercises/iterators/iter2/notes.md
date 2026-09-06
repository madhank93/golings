## iter2 — Seq2 yields pairs, in order

```go
func Enumerate(s []int) iter.Seq2[int, int] {
    return func(yield func(int, int) bool) {
        for i, v := range s {
            if !yield(i, v) { // index FIRST, value second
                return
            }
        }
    }
}
```

**Why it works**

- `iter.Seq2[K, V]` is `func(yield func(K, V) bool)` — two values per step, like
  ranging a map or a slice's index/value. `Enumerate` must yield the index as
  the first argument to match `for i, v := range Enumerate(s)`.

**Under the hood**

- The order is positional only. `Seq2[int, int]` cannot tell an index from a
  value, so the swapped version **compiles perfectly** and produces nonsense at
  the call site — the kind of bug only a test catches.

**Common mistake**

- Assuming a `Seq2` must be key/value. It is whatever pair the producer defines:
  index/element, key/value, or value/error — the last being how iterators report
  failures, since the protocol has no error channel of its own.

**Key detail:** when both halves have the same type, name them in the signature
(`func Enumerate(s []int) iter.Seq2[int, int]` with a doc comment saying which is
which) — the compiler will not do it for you.

**See also:** iter1 (`Seq`) · iter3 (`Seq2` → `Seq`) · range2 (map ranging) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — iter.Seq2: https://pkg.go.dev/iter#Seq2
- pkg.go.dev — maps.All: https://pkg.go.dev/maps#All
