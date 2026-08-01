## modern3 — range over a function (iter.Seq)

```go
func countUp(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := 1; i <= n; i++ {
            if !yield(i) { // consumer broke out early
                return
            }
        }
    }
}
for v := range countUp(3) { ... } // 1, 2, 3
```

**Why it works**

- Go 1.23 lets `for range` iterate a **function** of type `iter.Seq[V]` =
  `func(yield func(V) bool)`. Your function calls `yield(v)` for each value; the
  `range` loop body *is* that `yield`.

**Key detail:** always check `yield`'s return — `false` means the consumer did `break`
(or returned), so you must **stop** producing. This is the foundation of custom,
composable iterators without allocating a slice up front.

**References**

- The Go Blog — Range Over Function Types: https://go.dev/blog/range-functions
- pkg.go.dev — iter: https://pkg.go.dev/iter
