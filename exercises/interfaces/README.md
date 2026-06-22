# Interfaces

Interfaces are the heart of Go's polymorphism. An interface lists method
signatures; any type with those methods satisfies it **implicitly** — no
`implements` keyword. Interface values hold a (concrete type, value) pair.

- **interfaces1** — implicit satisfaction + polymorphism over `[]Shape`
- **interfaces2** — `fmt.Stringer` and the compile-time `var _ I = T{}` check
- **interfaces3** — recovering the concrete type with a type switch / assertion
- **interfaces4** — interface embedding (`ReadWriter` = `Reader` + `Writer`)

## Resources

- [A Tour of Go: Interfaces](https://go.dev/tour/methods/9)
- [Effective Go: Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go by Example: Interfaces](https://gobyexample.com/interfaces)
