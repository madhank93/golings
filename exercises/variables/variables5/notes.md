## variables5 — constants must be initialized

```go
const Pi = 3.14
```

**Why it works**

- `const Pi` with no value is illegal — unlike a `var`, a constant has **no zero
  value** to fall back on. A constant's value must be known at declaration.

**Nuance:** constants are fixed at **compile time**, so they can only be set from
constant expressions (literals, other constants, `iota`) — never from something
computed at run time like a function call.

**References**

- Go by Example — Constants: https://gobyexample.com/constants
- Go spec — Constants: https://go.dev/ref/spec#Constants
