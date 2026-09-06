## range1 — range a slice

```go
for _, v := range evenNumbers {
    fmt.Printf("%d is even\n", v)
}
```

**Why it works**

- Ranging a slice yields two values per iteration — the index and a copy of the
  element. This loop only needs the element, so the index is discarded with the
  blank identifier `_`.

**Common mistake**

- Writing `for v := range nums` when you wanted the values. With one variable
  you get the **index**, not the element — a silent logic bug on `[]int`, where
  both are ints and nothing fails to compile.

**Key detail:** all four forms are legal, and each says something different:
`for i, v := range s` (both), `for i := range s` (index), `for _, v := range s`
(value), and `for range s` (neither — just repeat `len(s)` times). Since Go 1.22
`for i := range 5` ranges an integer, no slice required.

**See also:** range2 (maps) · range3 (filtering) · slices4 (indexing by hand) ·
the [chapter](../README.md)

**References**

- Go spec — For statements with range clause: https://go.dev/ref/spec#For_range
- Go by Example — Range over Built-in Types: https://gobyexample.com/range-over-built-in-types
