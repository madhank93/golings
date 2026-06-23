# Generics

Generics (Go 1.18+) let functions and types work over a set of types via **type
parameters** with **constraints**: `func Map[T, U any](s []T, f func(T) U) []U`.
Constraints are interfaces (e.g. `any`, `comparable`, `constraints.Ordered`).

## Resources

- [Go blog: An Introduction to Generics](https://go.dev/blog/intro-generics)
- [A Tour of Go: Generics](https://go.dev/tour/generics/1)
- [Go by Example: Generics](https://gobyexample.com/generics)
