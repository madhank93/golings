## variables3 — a declaration needs a type or a value

```go
var x int
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `var x` alone gives Go nothing to infer from. Supply **either** a type
  (`var x int`) **or** an initializer (`var x = 0`). With the type only, `x`
  takes its **zero value** — `0` for `int`.

**Under the hood**

- Zero values are not a convention, they are a guarantee: allocated memory is
  zeroed before you see it. `0`, `false`, `""`, and `nil` for pointers, slices,
  maps, channels, funcs and interfaces. That is why `var buf bytes.Buffer` and
  `var mu sync.Mutex` are usable without a constructor.

**Common mistake**

- Reading the zero value as "unset". `0` and `""` are real values — a struct
  field holding them cannot tell you whether anyone assigned it. When that
  distinction matters, use a pointer, or the comma-ok form on maps.

**Key detail:** an unused local variable is a **compile error** in Go, not a
warning — so a declaration you never read has to go, or be assigned to `_`.

**See also:** variables1 · maps1 (comma-ok vs zero value) ·
primitive_types1 (bool's zero value) · the [chapter](../README.md)

**References**

- Go spec — The zero value: https://go.dev/ref/spec#The_zero_value
- Go by Example — Variables: https://gobyexample.com/variables
