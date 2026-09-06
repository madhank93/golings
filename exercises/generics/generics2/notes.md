## generics2 — constraints are type sets

```go
type Number interface {
    ~int | ~float64
}

func addNumbers[T Number](n1, n2 T) T {
    return n1 + n2
}
```

**Why it works**

- A constraint is an interface used as a **set of types**. `~int | ~float64`
  admits both, and `+` is legal in the body because every type in the set
  supports it. Widen the constraint to `any` and the same body stops compiling.

**Under the hood**

- `~int` means "any type whose **underlying type** is `int`", so a defined type
  like `type Celsius float64` satisfies `~float64`. Without the tilde only the
  exact predeclared type qualifies — and every domain type in the codebase is
  quietly excluded.

**Common mistake**

- Forgetting the return type. `func addNumbers[T Number](n1, n2 T)` promises
  nothing, so `return n1 + n2` has nowhere to go — the same
  `too many return values` as functions4, one abstraction level up.

**Key detail:** three constraints ship with Go: `any` (everything),
`comparable` (supports `==` — map keys), and `cmp.Ordered` (numbers and strings —
`<`, sorting, `min`/`max`). Reach for those before writing your own.

**See also:** generics1 · generics4 (two parameters) · typealias1 (why `~`
matters) · the [chapter](../README.md)

**References**

- Go spec — General interfaces (type sets): https://go.dev/ref/spec#General_interfaces
- pkg.go.dev — cmp.Ordered: https://pkg.go.dev/cmp#Ordered
