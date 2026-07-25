## enums1 — enums via iota

```go
type Weekday int

const (
    Sunday Weekday = iota // 0
    Monday                // 1
    Tuesday               // 2
    // ...
)
```

**Why it works**

- Go has **no `enum` keyword**. The idiom is a named integer type plus a `const`
  block using `iota`, which starts at `0` and increments by one per line.

**Nuance:** `iota` resets to `0` at each `const` block and counts **lines**, so
you only write it once — the type and `= iota` carry down to the following
constants implicitly. Skip a value with `_`, or scale it (`1 << iota`) for
bit-flag enums.

**References**

- Go by Example — Enums / Iota: https://gobyexample.com/enums
- Go spec — Iota: https://go.dev/ref/spec#Iota
