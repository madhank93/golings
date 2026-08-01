## generics3 — a generic data structure

```go
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(x T)        { s.items = append(s.items, x) }
func (s *Stack[T]) Pop() (T, bool)  { ... }
```

**Why it works**

- `Stack[T any]` is a generic type: one definition serves `Stack[int]`,
  `Stack[string]`, anything. Methods repeat the receiver's type parameter as
  `*Stack[T]` and use `T` for elements.

**Key detail:** `var zero T` yields `T`'s zero value — the only way to return "nothing"
generically when `Pop` fails, since you can't write `nil` or `0` for an unknown
`T`. Returning `(T, bool)` is the generic take on the comma-ok idiom.

**References**

- A Tour of Go — Generic types: https://go.dev/tour/generics/2
- The Go Blog — An Introduction to Generics: https://go.dev/blog/intro-generics
