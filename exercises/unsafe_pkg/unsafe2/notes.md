## unsafe2 — zero-copy []byte → string

```go
func bytesToString(b []byte) string {
    return unsafe.String(unsafe.SliceData(b), len(b))
}
```

**Why it works**

- `string(b)` allocates and copies, because a string is immutable and a slice is
  not — sharing would let one break the other's contract. `unsafe.String` builds
  a string header pointing straight at the slice's backing array, so the
  conversion costs nothing.

**Under the hood**

- `unsafe.SliceData(b)` returns the pointer to the first element and
  `unsafe.String(ptr, len)` wraps it. These builders (Go 1.17/1.20) replaced the
  old `*reflect.StringHeader` cast, which was never actually safe — the header
  types were not guaranteed to match the runtime's layout.

**Common mistake**

- Mutating `b` afterwards. The string's contents change with it, breaking an
  invariant that maps, comparisons, and the runtime all rely on — with no panic
  and no race report, just wrong answers later.

**Key detail:** the compiler already elides the copy in some common cases (a
`string(b)` used only as a map key, or in a `range`). Reach for this only after
profiling shows the conversion in a hot path, and keep it behind a small,
documented function — the way the standard library does.

**See also:** unsafe1 (safe layout queries) · strings2 (string internals) ·
pprof1 (proving it matters) · the [chapter](../README.md)

**References**

- pkg.go.dev — unsafe.String: https://pkg.go.dev/unsafe#String ·
  SliceData: https://pkg.go.dev/unsafe#SliceData
- Go 1.20 release notes: https://go.dev/doc/go1.20#unsafe
