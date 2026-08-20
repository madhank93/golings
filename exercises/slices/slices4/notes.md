## slices4 — stay inside the bounds

```go
names := []string{"John", "Maria", "Carl", "Peter"}
firstName := names[0]        // was names[3]
firstTwoNames := names[0:2]  // was names[5:10]
```

**Why it works**

- Valid indexes run `0 … len-1`, and slice bounds run `0 … cap`. `names[3]` is
  "Peter", not "John", and `names[5:10]` reaches past a 4-element backing array
  entirely.

**Under the hood**

- Indexing is checked against **length**, slicing against **capacity** — which is
  why `s[2:4]` can be legal on a slice whose `len` is 3 but whose `cap` is 5.
  Every access is bounds-checked; the compiler elides the check only when it can
  prove the index is safe.

**Common mistake**

- Expecting a silent empty result or a clamp, as some languages give. Go panics:
  `index out of range [4] with length 4`, or
  `slice bounds out of range [:10] with capacity 4`. A constant index the
  compiler can evaluate fails the build instead.

**Key detail:** the panic message names the offending value *and* the limit, so it
usually points straight at the off-by-one. `len(s)` before indexing, or `range`,
avoids the question altogether.

**See also:** arrays1 (zero-based indexing) · slices2 (bounds) ·
range1 (iterating without indexes) · the [chapter](../README.md)

**References**

- Go spec — Index expressions: https://go.dev/ref/spec#Index_expressions ·
  Slice expressions: https://go.dev/ref/spec#Slice_expressions
- Go by Example — Slices: https://gobyexample.com/slices
