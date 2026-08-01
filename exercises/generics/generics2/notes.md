## generics2 — constrain a type parameter

```go
type Number interface {
    int | float64 // a type SET: int OR float64
}

func addNumbers[T Number](n1, n2 T) T {
    return n1 + n2
}
```

**Why it works**

- `any` is too loose — you can't `+` arbitrary types. A constraint interface
  listing `int | float64` restricts `T` to types that support `+`, so `n1 + n2`
  compiles. The function then works for both `addNumbers(1, 2)` and
  `addNumbers(1.5, 2.5)`.

**Key detail:** constraint interfaces can list a **union of types** (a "type set"),
not just methods — that's the generics-era extension of `interface`. The standard
`golang.org/x/exp/constraints` package (and `cmp.Ordered`) provide ready-made
ones like `Ordered`.

**References**

- Go by Example — Generics: https://gobyexample.com/generics
- The Go Blog — An Introduction to Generics: https://go.dev/blog/intro-generics
