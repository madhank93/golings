## slices3 — append returns the new slice

```go
func addPeter(names []string) []string {
    return append(names, "Peter")
}
```

**Why it works**

- `append` takes the slice plus the values to add and **returns** the result. It
  cannot modify the caller's header in place, so the return value is the answer —
  ignoring it loses the append entirely.

**Under the hood**

- If the backing array has spare capacity, `append` writes into it and returns a
  header with `len+1` over the *same* array. If it does not, it allocates a
  bigger array, copies every element, and returns a header over the **new** one.
  That is why appending sometimes mutates a sibling slice and sometimes does not.

**Common mistake**

- `append(s, x)` as a statement, or `s2 := append(s, x)` and then assuming `s`
  changed too. Write `s = append(s, x)`. Inside a function, an append does not
  extend the caller's slice either — return it, as this exercise does.

**Key detail:** growth is amortised O(1) — the runtime roughly doubles small
slices — but each reallocation copies. When the final size is known,
`make([]T, 0, n)` removes every intermediate copy.

**See also:** slices1 (`len` vs `cap`) · slices2 (shared backing arrays) ·
morefn2 (spreading a slice into a variadic call) · the [chapter](../README.md)

**References**

- Go blog — the mechanics of 'append': https://go.dev/blog/slices
- src/runtime/slice.go — `growslice`: https://github.com/golang/go/blob/master/src/runtime/slice.go
