## stdlib3 — the slices package

```go
out := slices.Clone(nums)
slices.Sort(out)
slices.Reverse(out)
slices.Equal(got, want)
```

**Why it works**

- The `slices` package (Go 1.21+) provides generic helpers so you don't hand-roll
  sort/reverse/equal. `Clone` first, so sorting doesn't mutate the caller's slice.

**Nuance:** `slices.Sort` works on any ordered element type via generics — no
`sort.Interface` boilerplate needed for the common case. Remember that `Sort` and
`Reverse` mutate **in place**; `Clone` gives you an independent copy to protect
the input. Comparing slices needs `slices.Equal` (not `==`, which is illegal for
slices).

**References**

- pkg.go.dev — slices: https://pkg.go.dev/slices
