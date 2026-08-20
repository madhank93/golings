## variables5 — constants must be initialized

```go
const Pi = 3.14
```

**Why it works**

- `const Pi` with no value is illegal — unlike a `var`, a constant has **no zero
  value** to fall back on. A constant's value must be known at declaration.

**Under the hood**

- A constant is not stored anywhere; the compiler substitutes its value at every
  use. An *untyped* constant like this one also carries far more precision than
  any machine type while expressions are folded, and only takes a type when it
  is assigned or passed somewhere — which is why `Pi` works as a `float32` and
  a `float64` without conversion.

**Common mistake**

- `const start = time.Now()` — a function call is not a constant expression.
  Constants come from literals, other constants, `iota`, and operators over
  them. Anything computed at run time is a `var`.

**Key detail:** `const` is for values that are fixed **and** knowable at compile
time. That is a narrower set than "never changes at run time" — a lookup table
built at startup is still a `var`.

**See also:** variables6 (immutability) · enums1 (`iota`) ·
the [chapter](../README.md)

**References**

- Go spec — Constant declarations: https://go.dev/ref/spec#Constant_declarations
- Go blog — Constants: https://go.dev/blog/constants
