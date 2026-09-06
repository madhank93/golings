## generics1 — one function, any type

```go
func print[T any](value T) {
    fmt.Println(value)
}

print("Hello, World!")
print(42)
```

**Why it works**

- `[T any]` declares a **type parameter**. Each call instantiates the function
  for the argument's type, so one body serves a `string` and an `int` — with the
  compiler still checking every use.

**Under the hood**

- `T` is inferred from the argument, so the brackets are rarely written at a call
  site. Inference works from arguments only: a type parameter that appears just
  in the results has to be supplied explicitly (`Zero[int]()`).

**Common mistake**

- Reaching for `interface{}` (or `any`) as the parameter type instead. That
  compiles for this trivial body, but the value arrives type-erased — you cannot
  do arithmetic on it, compare it, or return it as its own type without an
  assertion. A type parameter keeps the type.

**Key detail:** `any` as a **constraint** and `any` as a **type** are different
jobs for the same word. As a constraint it means "no restriction"; the moment
the body needs an operator, the constraint must narrow to a type set that
supports it — generics2.

**See also:** generics2 (constraints) · generics3 (generic types) ·
interfaces3 (the `any` alternative) · the [chapter](../README.md)

**References**

- Go blog — An Introduction to Generics: https://go.dev/blog/intro-generics
- Go spec — Type parameter declarations: https://go.dev/ref/spec#Type_parameter_declarations
