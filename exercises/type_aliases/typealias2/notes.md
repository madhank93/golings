## typealias2 — a method on a defined type

```go
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", float64(c))
}
```

**Why it works**

- Because `Celsius` is a defined type you can attach methods to it. Implementing
  `String()` satisfies `fmt.Stringer`, so `fmt.Sprint(Celsius(25))` prints
  `"25°C"` automatically.

**Key detail — the recursion trap:** inside `String()` you must convert with
`float64(c)`. Writing `fmt.Sprintf("%g°C", c)` (or `fmt.Sprint(c)`) would call
`c.String()` **again**, looping forever until the stack overflows. Convert to the
underlying type before formatting.

**References**

- pkg.go.dev — fmt.Stringer: https://pkg.go.dev/fmt#Stringer
