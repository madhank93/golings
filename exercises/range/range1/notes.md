## range1 — iterate a slice with range

```go
for _, v := range evenNumbers {
    fmt.Printf("%d is even\n", v)
}
```

**Why it works**

- `range` over a slice yields two values each step: the **index** and a **copy of
  the element**. Here the index is unused, so `_` discards it and `v` is the
  value.

**Nuance:** `v` is a **copy** — assigning to it doesn't change the slice. To
mutate elements in place, index through them: `for i := range s { s[i] = ... }`.

**References**

- Go by Example — Range: https://gobyexample.com/range
