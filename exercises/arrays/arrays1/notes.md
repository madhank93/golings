## arrays1 — arrays are zero-indexed

```go
func first(colors [3]string) string { return colors[0] }
func last(colors [3]string) string  { return colors[2] }
```

**Why it works**

- The first element lives at index `0`; the last element of a length-3 array is
  at index `len-1 = 2`. The broken code left the index blank.

**Nuance:** an array's length is **part of its type** — `[3]string` and
`[4]string` are different types, and the size is fixed at compile time. Indexing
out of range (`colors[3]`) is a **run-time panic**, not a compile error, so the
bounds are yours to get right.

**References**

- Go by Example — Arrays: https://gobyexample.com/arrays
