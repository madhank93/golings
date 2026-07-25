## variables3 — a declaration needs a type or a value

```go
var x int
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `var x` alone gives Go nothing to infer from. Supply **either** a type
  (`var x int`) **or** an initializer (`var x = 0`). With the type only, `x`
  takes its **zero value** — `0` for `int`.

**Nuance:** every Go type has a zero value (`0`, `false`, `""`, `nil`), so a
declared-but-uninitialized variable is always usable, never garbage.

**References**

- Go by Example — Variables: https://gobyexample.com/variables
