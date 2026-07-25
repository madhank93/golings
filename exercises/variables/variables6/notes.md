## variables6 — constants can't be reassigned

```go
var x = 10 // was: const x = 10
func main() {
    fmt.Println(x)
    x = x + 1 // now legal — x is a variable
    fmt.Println(x)
}
```

**Why it works**

- The code reassigns `x` with `x = x + 1`. A `const` is immutable, so that line
  won't compile. Changing the declaration to `var` makes `x` mutable.

**Nuance:** the fix is about **choosing the right kind of binding**. Reach for
`const` only when a value never changes and is known at compile time; anything
you reassign must be a `var`.

**References**

- Go by Example — Constants: https://gobyexample.com/constants
