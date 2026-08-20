## enums2 — give the constants names

```go
func (c Color) String() string {
    switch c {
    case Red:
        return "Red"
    case Green:
        return "Green"
    case Blue:
        return "Blue"
    }
    return fmt.Sprintf("Color(%d)", int(c))
}
```

**Why it works**

- `String() string` satisfies `fmt.Stringer`, so `fmt.Println(Red)` prints `Red`
  rather than `0`. The method is on the named type, which is the reason to
  declare `type Color int` in the first place.

**Under the hood**

- `fmt` type-asserts each argument against `Stringer` and calls it when present.
  Use a **value receiver**: with `func (c *Color) String()`, printing a `Color`
  value would not use it, since the method would belong to `*Color` only.

**Common mistake**

- Returning `""` for an unhandled value, or formatting the receiver with `%v` in
  the fallback. The first hides bugs; the second calls `String()` again and
  recurses until the stack overflows. Convert — `int(c)` — in the fallback.

**Key detail:** the compiler does **not** check exhaustiveness. Adding a fourth
colour will not break this switch or any other; the `exhaustive` linter adds
that check, and `stringer` (`//go:generate stringer -type=Color`) writes the
method for you once the list grows.

**See also:** enums1 (`iota`) · interfaces2 (`Stringer`) · typealias2 (the same
method on a float type) · the [chapter](../README.md)

**References**

- pkg.go.dev — fmt.Stringer: https://pkg.go.dev/fmt#Stringer
- pkg.go.dev — stringer: https://pkg.go.dev/golang.org/x/tools/cmd/stringer
