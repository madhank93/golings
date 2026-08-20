## methods2 — methods on any named type

```go
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```

**Why it works**

- A receiver can be **any type defined in the package**, not just a struct.
  `Celsius` has `float64` as its underlying type, so converting the receiver
  gives ordinary arithmetic back.

**Under the hood**

- The conversion is required because `Celsius` is a **distinct type**, not an
  alias — same representation, different identity, no implicit mixing. The
  conversion itself generates no code; it only tells the compiler you meant it.

**Common mistake**

- Skipping the conversion and writing `c*9/5 + 32`. That actually compiles (the
  untyped constants adopt `Celsius`) but the result is a `Celsius`, which will
  not satisfy a `float64` return — the error appears at the `return`, one step
  away from the cause.

**Key detail:** this pattern is everywhere in the standard library —
`time.Duration` is a named `int64` with methods, and `http.HandlerFunc` is a
named *func* type with a method. You cannot attach methods to a type from
another package; define your own type around it first.

**See also:** methods1 (receiver choice) · typealias1 (why the conversion is
needed) · typealias2 (`String()` on the same type) · the [chapter](../README.md)

**References**

- Go spec — Method declarations: https://go.dev/ref/spec#Method_declarations
- pkg.go.dev — time.Duration: https://pkg.go.dev/time#Duration
