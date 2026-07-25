## primitive_types5 — use Go's real numeric type names

```go
var n1 int = 101
var n2 float64 = 0.99
```

**Why it works**

- `integer` and `float` are not Go types. The real names are `int` (and sized
  variants `int8/16/32/64`) and `float64` (or `float32`).

**Nuance:** `int` is not a fixed width — it's platform-sized (32- or 64-bit).
When you need an exact width use the sized names. And Go never converts between
numeric types implicitly: mixing an `int` and a `float64` in arithmetic is a
compile error until you convert one explicitly.

**References**

- A Tour of Go — Basic types: https://go.dev/tour/basics/11
- Go spec — Numeric types: https://go.dev/ref/spec#Numeric_types
