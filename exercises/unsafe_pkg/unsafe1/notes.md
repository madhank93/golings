## unsafe1 — Offsetof reports layout

```go
func offsetOfC() uintptr {
    return unsafe.Offsetof(Record{}.C)
}
```

**Why it works**

- `unsafe.Offsetof` reports the byte offset of a field within its struct. It is a
  **constant expression** evaluated at compile time — no reflection, no
  allocation, no run-time cost at all.

**Under the hood**

- For `struct{ A byte; B int64; C int32 }` the answer is 16, not 9: `B` needs
  8-byte alignment, so seven padding bytes follow `A`. The compiler **does not
  reorder fields**, so declaration order determines size — reordering to
  `B, C, A` takes this struct from 24 bytes to 16.

**Common mistake**

- Assuming `unsafe.Sizeof` includes the data a field points at. It does not: for
  a slice field it is the 24-byte header, not the elements. Same for strings,
  maps, and pointers.

**Key detail:** this is the one corner of `unsafe` with no danger — the three
layout functions (`Sizeof`, `Offsetof`, `Alignof`) only compute numbers. Use them
with `go vet`'s `fieldalignment` check when a struct exists in millions of copies;
ignore them otherwise.

**See also:** unsafe2 (the risky idiom) · structs1 (field layout) ·
pprof2 (measuring memory) · the [chapter](../README.md)

**References**

- pkg.go.dev — unsafe.Offsetof: https://pkg.go.dev/unsafe#Offsetof
- Go spec — Package unsafe: https://go.dev/ref/spec#Package_unsafe
