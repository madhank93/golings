# Enums (iota)

Go has no `enum` keyword. The idiom is a named integer type plus a `const`
block using `iota`, which produces 0, 1, 2, … for successive constants. Add a
`String()` method for readable names.

- **enums1** — `iota` constant generation
- **enums2** — `iota` + `String()` for named values

## Resources

- [Go by Example: Iota](https://gobyexample.com/enums)
- [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
- [Go spec: Iota](https://go.dev/ref/spec#Iota)
