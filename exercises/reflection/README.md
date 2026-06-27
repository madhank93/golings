# Reflection

The `reflect` package lets a program inspect and manipulate values whose types
aren't known at compile time. Reach for it sparingly — most Go code prefers
interfaces and generics — but it powers encoders, ORMs, and validation.

- **reflect1** — read a value's `reflect.Kind` with `reflect.TypeOf`
- **reflect2** — walk a struct's fields with `reflect.ValueOf` and act on the string ones

## Resources

- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [reflect package](https://pkg.go.dev/reflect)
