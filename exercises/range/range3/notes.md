## range3 — filter with range + append

```go
evenNumbers := []int{}
for _, n := range numbers {
    if n%2 == 0 {
        evenNumbers = append(evenNumbers, n)
    }
}
```

**Why it works**

- Range each number, test `n%2 == 0`, and append the ones that pass. Starting
  from an empty slice and appending is the standard "filter" pattern.

**Key detail:** `evenNumbers = append(...)` reassigns each iteration — the growing
slice must be captured back into the variable. Starting from `[]int{}` (not a
`nil` slice) makes the result a non-nil empty slice when nothing matches, which
`reflect.DeepEqual` treats as distinct from `nil`.

**References**

- Go by Example — Range: https://gobyexample.com/range
- Go by Example — Slices: https://gobyexample.com/slices
