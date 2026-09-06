## applied1 — implementing sort.Interface

```go
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

sort.Sort(ByAge(people))
```

**Why it works**

- `sort.Sort` needs three methods and knows nothing else about the data.
  Implementing them on a **defined slice type** lets it order any collection —
  implicit interface satisfaction doing the work.

**Under the hood**

- The receivers are **values**, and `Swap` still mutates the caller's data: `a` is
  a slice, so the copy shares the same backing array. That is the "copied header,
  shared data" rule from the pointers chapter being useful rather than surprising.

**Common mistake**

- Writing `<=` in `Less`. The contract requires a strict weak ordering —
  `Less(i, i)` must be false — and a non-strict comparison can make the sort
  behave erratically rather than merely order oddly.

**Key detail:** modern Go usually writes
`slices.SortFunc(people, func(a, b Person) int { return cmp.Compare(a.Age, b.Age) })`
— no type declaration, no interface dispatch per comparison. Note the flipped
convention: `sort.Interface` wants a `bool`, `slices.SortFunc` a three-way `int`.
Neither is stable; `sort.SliceStable` and `slices.SortStableFunc` are.

**See also:** methods2 (methods on non-struct types) · interfaces1 ·
stdlib3 (`slices`) · applied2 · the [chapter](../README.md)

**References**

- pkg.go.dev — sort.Interface: https://pkg.go.dev/sort#Interface ·
  slices.SortFunc: https://pkg.go.dev/slices#SortFunc
