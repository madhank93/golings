## slices1 — make needs the element type

```go
a := make([]int, 3, 10) // len 3, cap 10
fmt.Println(len(a), cap(a)) // 3 10
```

**Why it works**

- `make` builds a slice from three parts: the **type** (`[]int`), the length, and
  an optional capacity. Leaving the type out gives it nothing to allocate.

**Under the hood**

- The result is a three-word header — pointer, length, capacity — over a backing
  array of 10 zeroed ints. `len` is what you may index (`a[3]` panics); `cap` is
  how far `append` can grow before it must allocate a new array and copy.

**Common mistake**

- Confusing the two arguments. `make([]int, 10)` gives ten **zero elements**, not
  an empty slice with room for ten. For "empty but preallocated" the form is
  `make([]int, 0, 10)` — the one that makes an append loop allocate exactly once.

**Key detail:** the zero value of a slice is `nil`, and it is usable —
`len`, `range`, and `append` all work on it. Prefer `var s []int` to
`[]int{}` for an empty slice.

**See also:** slices3 (`append` and growth) · slices4 (bounds) ·
maps1 (`make` for maps) · the [chapter](../README.md)

**References**

- Go blog — Go Slices: usage and internals: https://go.dev/blog/slices-intro
- Go spec — Making slices, maps and channels: https://go.dev/ref/spec#Making_slices_maps_and_channels
