## slices2 — sub-slice with [low:high]

```go
func lastTwo(names [4]string) []string {
    return names[2:4]
}
```

**Why it works**

- `names[2:4]` takes elements at indexes 2 and 3 — the last two. The bound is
  **half-open**: `low` is included, `high` is **not**, so `[2:4]` gives 2 items,
  not 3.

**Key detail:** slicing does **not** copy — the result shares the same backing array
as the original. Writing through the sub-slice can change the parent (and vice
versa) until an `append` forces a re-allocation.

**References**

- Go by Example — Slices: https://gobyexample.com/slices
