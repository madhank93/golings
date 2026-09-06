## generics4 — two type parameters

```go
func Reduce[A, B any](items []A, init B, f func(B, A) B) B {
    acc := init
    for _, v := range items {
        acc = f(acc, v)
    }
    return acc
}
```

**Why it works**

- `A` is the element type and `B` the accumulator, and they are **independent** —
  which is what lets one function fold a `[]int` into an `int` and a `[]string`
  into a `string`. The fold itself is a plain loop over `f`.

**Under the hood**

- Both parameters are inferred from the call: `A` from the slice, `B` from
  `init`. That is also why the accumulator's type is pinned by the *initial
  value* — `Reduce(nums, 0, …)` folds into an `int`, `Reduce(nums, 0.0, …)`
  into a `float64`.

**Common mistake**

- Getting `f`'s parameter order backwards. The signature is
  `func(B, A) B` — accumulator first, element second — matching the order the
  loop applies them. A swapped literal fails to compile, which is the type system
  earning its keep.

**Key detail:** this is the shape most generic code takes in practice: a function
that takes a function. It is also the shape to be careful with — Go's advice is
to write the concrete version first and generalise only when you have written it
three times.

**See also:** generics2 (constraints) · generics3 (generic types) ·
anonymous_functions2 (function types) · iter1 (the lazy equivalent) ·
the [chapter](../README.md)

**References**

- Go blog — When To Use Generics: https://go.dev/blog/when-generics
- Go spec — Type inference: https://go.dev/ref/spec#Type_inference
