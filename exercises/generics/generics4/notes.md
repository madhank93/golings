## generics4 — a generic Reduce

```go
func Reduce[A, B any](items []A, init B, f func(B, A) B) B {
    acc := init
    for _, item := range items {
        acc = f(acc, item)
    }
    return acc
}
```

**Why it works**

- Two type parameters: `A` (element type) and `B` (accumulator type). `Reduce`
  folds `[]A` into a single `B` by repeatedly applying `f`. Summing ints uses
  `A=int, B=int`; concatenating strings uses `A=string, B=string` — both inferred.

**Nuance:** separating `A` and `B` lets the result type differ from the element
type — e.g. reduce `[]string` into an `int` length. This is the generic version
of fold/reduce; higher-order functions plus type parameters give you reusable,
type-safe building blocks.

**References**

- The Go Blog — An Introduction to Generics: https://go.dev/blog/intro-generics
- Go by Example — Generics: https://gobyexample.com/generics
