## unsafe1 — unsafe.Offsetof

```go
func offsetOfC() uintptr {
    return unsafe.Offsetof(Record{}.C)
}
```

**Why it works**

- `unsafe.Offsetof(x.Field)` reports the **byte offset** of `Field` within its
  struct — compile-time layout introspection with zero runtime cost and no
  reflection.

**Key detail:** the offset isn't just the sum of field sizes — the compiler inserts
**padding** to satisfy alignment (`A byte` then `B int64` leaves 7 bytes of
padding so `B` lands on an 8-byte boundary). That's why field **order** affects
struct size. `unsafe` is for interop/serialization/hot paths; it bypasses type
safety, so use it sparingly and deliberately.

**References**

- pkg.go.dev — unsafe.Offsetof: https://pkg.go.dev/unsafe#Offsetof
- Go spec — Package unsafe: https://go.dev/ref/spec#Package_unsafe
