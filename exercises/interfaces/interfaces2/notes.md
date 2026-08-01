## interfaces2 — implement fmt.Stringer

```go
func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
var _ fmt.Stringer = Point{} // compile-time proof Point satisfies Stringer
```

**Why it works**

- `fmt` looks for a `String() string` method; implementing it makes `%v`/`Println`
  print `(3, 4)` instead of the default `{3 4}`.

**Key detail:** the line `var _ fmt.Stringer = Point{}` is a common idiom — a
throwaway assignment that makes the **compiler verify** `Point` satisfies
`Stringer`. If the method is missing or has the wrong signature, the build fails
at that line with a clear message, instead of silently falling back to default
formatting.

**References**

- fmt.Stringer: https://pkg.go.dev/fmt#Stringer
- Go by Example — Interfaces: https://gobyexample.com/interfaces
