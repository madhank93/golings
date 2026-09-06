## generics3 — a generic data structure

```go
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(x T) { s.items = append(s.items, x) }

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top, true
}
```

**Why it works**

- `Stack[T any]` is a generic type: one definition serves `Stack[int]`,
  `Stack[string]`, anything. Methods repeat the receiver's type parameter as
  `*Stack[T]` and use `T` for elements.

**Under the hood**

- `var zero T` is the only way to produce "nothing" for an unknown type — you
  cannot write `nil` or `0` for a `T` that might be either. Returning `(T, bool)`
  is the generic spelling of the comma-ok idiom you already use on maps.

**Common mistake**

- Trying to give a method its own type parameter —
  `func (s *Stack[T]) Map[U any](...)`. Methods may only use the **type's**
  parameters; anything else has to be a plain function taking the stack as an
  argument.

**Key detail:** the pointer receiver is required for `Push` and `Pop` because both
mutate `items`; a value receiver would append to a copy. Same rule as any other
type — generics change nothing about method sets.

**See also:** generics4 (function parameters) · methods1 (receivers) ·
slices3 (`append` semantics) · maps3 (comma-ok) · the [chapter](../README.md)

**References**

- Go blog — An Introduction to Generics: https://go.dev/blog/intro-generics
- A Tour of Go — Generic types: https://go.dev/tour/generics/2
