## range2 — range a map

```go
for name, phone := range phoneBook {
    fmt.Printf("%s has the %s phone\n", name, phone)
}
```

**Why it works**

- Over a map, `range` yields **key and value** rather than index and element. One
  variable gives just the key: `for name := range phoneBook`.

**Under the hood**

- Iteration starts at a **randomly chosen bucket and offset**, so the order
  differs on every run of the same program. That randomisation was added
  deliberately, so that nobody can depend on an order the implementation is free
  to change.

**Common mistake**

- Printing a map in a test and asserting on the output. It passes locally and
  fails eventually. Collect the keys and sort them when order matters:

```go
for _, k := range slices.Sorted(maps.Keys(m)) { … }   // Go 1.23
```

**Key detail:** deleting an entry during iteration is safe — an entry not yet
reached simply is not produced. Adding during iteration is unspecified: the new
entry may or may not be visited.

**See also:** range1 (slices) · maps3 (comma-ok) · mapspkg1 (`maps.Keys`) ·
the [chapter](../README.md)

**References**

- Go spec — For statements with range clause: https://go.dev/ref/spec#For_range
- Go blog — Go maps in action: https://go.dev/blog/maps
