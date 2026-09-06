# Generics

Before Go 1.18, writing one function that worked for both `[]int` and
`[]string` meant either copy-pasting it per type or routing everything through
`interface{}` and giving up type checking at the door. Generics remove that
trade: a function or type can take **type parameters**, the compiler checks each
use against a **constraint**, and the code stays as statically typed as anything
you write by hand.

The four exercises build the whole vocabulary — a type parameter, a constraint
that admits several types, a generic data structure, and a generic function that
takes a function.

## 1. Type parameters

```go
func Print[T any](value T) {
    fmt.Println(value)
}

Print("Hello")   // T inferred as string
Print(42)        // T inferred as int
Print[float64](3)  // explicit, rarely needed
```

The square brackets after the name declare the parameters; each needs a
constraint, and `any` means "no restriction". `generics1` is a function missing
its parameter type — the pre-generics answer would have been `value interface{}`,
which prints the same and tells the compiler nothing.

Inference means you almost never write the brackets at a call site. It works
from the *arguments*, so a type parameter used only in the return type must be
supplied explicitly.

## 2. Constraints are interfaces

A constraint is an interface used as a type set rather than a method set:

```go
type Number interface {
    ~int | ~int64 | ~float64      // a union of type terms
}

func addNumbers[T Number](a, b T) T {
    return a + b       // legal: every type in the set supports +
}
```

That is `generics2`. The constraint is what makes the body checkable — `+` is
allowed because it works for every type in the set. Widen the set to `any` and
the same body stops compiling.

The `~` matters: `~int` means "any type whose **underlying type** is `int`", so a
defined type like `type Celsius float64` satisfies `~float64`. Without the tilde
only the exact type qualifies, and every domain type in your codebase is
excluded.

Three constraints come with the language or the standard library:

| Constraint | Admits | Enables |
|---|---|---|
| `any` | everything | assignment, passing around |
| `comparable` | types supporting `==` | map keys, equality |
| [`cmp.Ordered`](https://pkg.go.dev/cmp#Ordered) | numbers and strings | `<`, `>`, sorting, `min`/`max` |

`cmp.Ordered` (Go 1.21) replaced the older `golang.org/x/exp/constraints.Ordered`
for most uses.

## 3. Generic types

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(x T)      { s.items = append(s.items, x) }
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T          // the only way to say "nothing" generically
        return zero, false
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top, true
}
```

`generics3`. Two details that trip people up: the **receiver repeats the type
parameter** (`*Stack[T]`, not `*Stack`), and **methods cannot introduce new type
parameters** — only the type's own. A method that needs its own parameter has to
be a plain function.

`var zero T` is the generic zero value. You cannot write `nil` or `0` for an
unknown `T`, so returning `(T, bool)` is the generic spelling of the comma-ok
idiom.

## 4. Multiple parameters, and functions as parameters

```go
func Reduce[A, B any](items []A, init B, f func(B, A) B) B {
    acc := init
    for _, v := range items {
        acc = f(acc, v)
    }
    return acc
}

Reduce([]int{1,2,3,4}, 0, func(acc, n int) int { return acc + n })      // 10
Reduce([]string{"a","b"}, "", func(acc, s string) string { return acc+s }) // "ab"
```

`generics4`. Two independent parameters — the element type and the accumulator
type — let one function fold a `[]int` into an `int` and a `[]string` into a
`string`. Inference fills both in from the arguments.

## 5. When *not* to reach for them

The Go team's own guidance is to write the concrete version first. Generics earn
their place when:

- you are writing a **container** (`Stack[T]`, `Set[T]`, a cache),
- you are writing a **general algorithm over slices, maps or channels** —
  which is what the `slices` and `maps` packages now are,
- or you would otherwise write the same function three times.

They do **not** help when the types need genuinely different behaviour — that is
what an interface is for — and a generic signature is always harder to read than
a concrete one. `any` in a signature where a type parameter belongs is a code
smell; a type parameter where a plain `[]string` belongs is the opposite one.

## Gotchas

- **Methods cannot have their own type parameters.** Make it a function.
- **The receiver must repeat the parameter**: `func (s *Stack[T]) …`.
- **Forgetting `~`** in a union excludes every defined type built on those
  underlying types.
- **`comparable` is not `cmp.Ordered`** — it gives you `==`, not `<`.
- **Inference works from arguments only**; a parameter that appears only in the
  results must be given explicitly.
- **`var zero T`** is how you produce a zero value you cannot name.
- **Generic code is compiled per type shape (GC shape stenciling)**, so a
  generic function is not free — but the cost is compile time and code size, not
  boxing on every call.

## The exercises

- **generics1** — give the function a type parameter so it accepts a string and
  an int.
- **generics2** — widen the constraint to a union of numeric types and give the
  function its parameter and return type.
- **generics3** — implement `Push`/`Pop` on a generic `Stack[T]`, including the
  generic zero value.
- **generics4** — fold any slice into any accumulator with two type parameters.

## Source references

- [Go blog: An Introduction to Generics](https://go.dev/blog/intro-generics) ·
  [When to use generics](https://go.dev/blog/when-generics)
- [Go spec: Type parameters](https://go.dev/ref/spec#Type_parameter_declarations) ·
  [Type constraints](https://go.dev/ref/spec#General_interfaces)
- [pkg.go.dev: cmp.Ordered](https://pkg.go.dev/cmp#Ordered) ·
  [slices](https://pkg.go.dev/slices) · [maps](https://pkg.go.dev/maps)
- [A Tour of Go: Generics](https://go.dev/tour/generics/1)

**Next: [modern](../modern/) →** — the built-ins and loop semantics that landed
alongside generics and changed how everyday Go is written.
