## stdlib3 — the slices package

```go
func sortDesc(nums []int) []int {
    out := slices.Clone(nums)
    slices.Sort(out)
    slices.Reverse(out)
    return out
}
```

**Why it works**

- `slices.Sort` orders ascending and `slices.Reverse` flips it — or, in one call,
  `slices.SortFunc(out, func(a, b int) int { return cmp.Compare(b, a) })`. The
  `Clone` first is what keeps the caller's slice untouched.

**Under the hood**

- `SortFunc`'s comparison returns an **int** (negative, zero, positive), not the
  `bool` that `sort.Slice` wanted. The three-way result matches `cmp.Compare`
  and lets the sort avoid a second comparison per pair.

**Common mistake**

- Sorting in place and returning it, thinking the input is safe. `slices.Sort`
  mutates the slice it is given, and a slice argument shares the caller's backing
  array — so without the `Clone` the caller's data is reordered too.

**Key detail:** `slices.Sort` is **not stable**; `slices.SortStableFunc` is. The
package also carries `Contains`, `Index`, `Equal`, `BinarySearch`, `Max`/`Min`,
`Compact` and `Insert` — all generic, all replacing loops people used to write
by hand.

**See also:** slices3 (`append` and headers) · mapspkg1 (the same copy trap for
maps) · generics2 (`cmp.Ordered`) · the [chapter](../README.md)

**References**

- pkg.go.dev — slices: https://pkg.go.dev/slices
- pkg.go.dev — cmp.Compare: https://pkg.go.dev/cmp#Compare
