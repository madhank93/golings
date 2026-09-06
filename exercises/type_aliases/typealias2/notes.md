## typealias2 — a defined type can carry methods

```go
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", float64(c))
}

fmt.Sprint(Celsius(25)) // "25°C"
```

**Why it works**

- Methods may be declared on any type **defined in this package**, so `Celsius`
  gets a `String()` and thereby satisfies `fmt.Stringer` — which is what makes
  `fmt.Sprint` print `25°C`.

**Under the hood**

- `float64(c)` inside is not decoration. Formatting the receiver itself
  (`fmt.Sprintf("%v", c)`) calls `String()` again, which formats `c` again — an
  infinite recursion that ends in `stack overflow`. Converting to the underlying
  type breaks the cycle.

**Common mistake**

- Trying the same on an alias. `type MyTime = time.Time` names *time's* type, and
  methods can only be declared on types defined locally — so an alias gives you
  nowhere to attach behaviour.

**Key detail:** attaching methods is the main reason to define a type rather than
alias one. `time.Duration` (a named `int64`) and `http.HandlerFunc` (a named
func type) are the standard library doing exactly this.

**See also:** typealias1 (distinctness) · typealias3 (aliases) ·
interfaces2 (`Stringer`) · enums2 · the [chapter](../README.md)

**References**

- Go spec — Method declarations: https://go.dev/ref/spec#Method_declarations
- pkg.go.dev — fmt.Stringer: https://pkg.go.dev/fmt#Stringer
