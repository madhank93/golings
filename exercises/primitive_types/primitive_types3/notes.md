## primitive_types3 — every format argument must exist

```go
who := "Gopher"
country := "India"
fmt.Printf("Hello, I am %s and live in %s\n", who, country)
```

**Why it works**

- The format string has two `%s` verbs, so `Printf` needs two declared
  arguments. Declaring both `who` and `country` satisfies them.

**Common mistake**

- Trusting the count by eye. `Printf` is variadic (`a ...any`), so a wrong count
  is a *runtime* artefact, not a build failure: too few prints
  `%!s(MISSING)`, too many appends `%!(EXTRA string=India)`.

**Key detail:** verbs are filled **left to right**, one argument each, so order
matters as much as count. `Println` needs none of this — it takes values
directly and puts spaces between them, which is often the better choice for
quick output.

**See also:** primitive_types2 (verbs) · the [chapter](../README.md)

**References**

- pkg.go.dev — fmt: https://pkg.go.dev/fmt
- Go by Example — String Formatting: https://gobyexample.com/string-formatting
