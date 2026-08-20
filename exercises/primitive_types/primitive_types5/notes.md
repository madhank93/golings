## primitive_types5 — use Go's real numeric type names

```go
var n1 int = 101
var n2 float64 = 0.99
```

**Why it works**

- `integer` and `float` are not Go types. The real names are `int` (with sized
  variants `int8/16/32/64`), the unsigned `uint…` family, and `float32` /
  `float64`.

**Under the hood**

- `int` is platform-sized (64-bit on modern targets) and is a **distinct type**
  from `int64` even where the widths match — passing one where the other is
  expected still needs a conversion. Go never converts numeric types
  implicitly, and `T(v)` conversions truncate silently: `byte(300)` is 44.

**Common mistake**

- Reaching for sized types everywhere. `int` is right for loop counters,
  lengths, and IDs; the sized names are for wire formats, file layouts, and
  places where the width is part of the contract.

**Key detail:** untyped constants are the exception that makes the strictness
bearable — `var f float64 = 3` works because `3` has no type yet. Once a value
lives in a variable its type is fixed and conversions become explicit.

**See also:** primitive_types4 (`byte`/`rune`) · typealias1 (defined types vs
aliases) · the [chapter](../README.md)

**References**

- Go spec — Numeric types: https://go.dev/ref/spec#Numeric_types ·
  Conversions: https://go.dev/ref/spec#Conversions
- Go blog — Constants: https://go.dev/blog/constants
