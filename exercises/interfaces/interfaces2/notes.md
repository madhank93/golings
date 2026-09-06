## interfaces2 — implement fmt.Stringer

```go
func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

var _ fmt.Stringer = Point{} // compile-time satisfaction check
```

**Why it works**

- `fmt.Stringer` is `interface{ String() string }`. Once `Point` has that method,
  every `fmt` verb that prints a value — `%v`, `Print`, `Sprintf` — routes
  through it, so the type controls its own presentation.

**Under the hood**

- `fmt` type-asserts each argument against `Stringer` (and `error`, and
  `Formatter`) at run time and calls whichever it finds. Nothing is registered
  in advance; the check is on the interface's type word.

**Common mistake**

- Formatting the receiver itself inside `String()`. `fmt.Sprintf("%v", p)` calls
  `String()` again and recurses until the stack dies. Format the **fields**, as
  here, or convert to the underlying type first.

**Key detail:** `var _ fmt.Stringer = Point{}` costs nothing at run time and turns
a future refactor into a compile error. Note the receiver form matters: with
`func (p *Point) String()`, the assertion needs `&Point{}` and
`fmt.Println(p)` on a *value* would not use it.

**See also:** interfaces1 · typealias2 (the same method on a named float) ·
enums2 (`String()` for constants) · methods1 (method sets) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — fmt.Stringer: https://pkg.go.dev/fmt#Stringer
- Go spec — Interface types: https://go.dev/ref/spec#Interface_types
