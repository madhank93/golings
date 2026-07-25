## unsafe2 — zero-copy []byte → string

```go
func bytesToString(b []byte) string {
    return unsafe.String(unsafe.SliceData(b), len(b))
}
```

**Why it works**

- `string(b)` **allocates** a new array and copies the bytes. `unsafe.String(ptr,
  len)` builds a string header that **shares** `b`'s backing array — no
  allocation, no copy. `unsafe.SliceData(b)` gets the pointer to that array.

**Nuance:** the shared string and the slice now alias the same memory, so you
**must not mutate `b` afterwards** — strings are supposed to be immutable, and
writing through `b` would violate that (and can corrupt map keys, etc.). This is
a hot-path optimization only; reach for it when profiling proves the copy matters.

**References**

- pkg.go.dev — unsafe.String: https://pkg.go.dev/unsafe#String
