## applied1 — sort.Interface

```go
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
sort.Sort(ByAge(people))
```

**Why it works**

- `sort.Sort` orders **anything** that implements `Len`, `Less`, and `Swap`.
  `Less` defines the ordering (here, ascending age), so `sort.Sort` rearranges the
  slice accordingly.

**Nuance:** the named slice type `ByAge` exists purely to hang the three methods
on `[]Person` — swap in a `ByName` type for a different order. For simple cases
`sort.Slice(people, func(i, j int) bool { ... })` skips the boilerplate, but
`sort.Interface` is the reusable, composable version.

**References**

- pkg.go.dev — sort.Interface: https://pkg.go.dev/sort#Interface
