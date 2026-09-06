## typealias1 — a defined type is a new type

```go
type Celsius float64

var boiling Celsius = 100
fmt.Printf("%.0f°F\n", toFahrenheit(float64(boiling))) // explicit conversion
```

**Why it works**

- `Celsius` and `float64` share a representation but are **different types**, so
  passing one where the other is expected fails. `float64(boiling)` converts
  explicitly at the call site.

**Under the hood**

- The conversion generates **no code** — the bits are identical. It exists purely
  so the compiler can check intent, which is the whole trade: zero run-time cost,
  visible in the source.

**Common mistake**

- Expecting the new type to inherit methods. A defined type starts with an
  **empty method set**: `type MyTime time.Time` has none of `time.Time`'s
  methods, only its layout.

**Key detail:** this is how you stop unit mix-ups at compile time — `UserID` vs
`OrderID`, `Celsius` vs `Fahrenheit`, all `int`s and `float64`s underneath. The
gap: untyped constants still convert freely, so `var c Celsius = 100` needs no
conversion and a literal mistake still slips through.

**See also:** typealias2 (methods on the type) · typealias3 (a true alias) ·
methods2 · enums1 · the [chapter](../README.md)

**References**

- Go spec — Type definitions: https://go.dev/ref/spec#Type_definitions
- Go spec — Conversions: https://go.dev/ref/spec#Conversions
