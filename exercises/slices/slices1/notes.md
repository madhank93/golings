## slices1 — make a slice with an element type

```go
a := make([]int, 3, 10) // len 3, cap 10
```

**Why it works**

- `make(, 3, 10)` is missing the type. `make([]int, 3, 10)` builds an `[]int`
  with **length** 3 (usable elements, all zero) and **capacity** 10 (room before
  a re-allocation is needed).

**Nuance:** length vs capacity — `len` is how many elements exist now, `cap` is
how many the backing array can hold before `append` must grow it. A slice is a
lightweight view (pointer + len + cap) over an underlying array.

**References**

- Go by Example — Slices: https://gobyexample.com/slices
- The Go Blog — Slices: usage and internals: https://go.dev/blog/slices-intro
