## variables1 — a variable needs a name

```go
var x = 5
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `var = 5` is a syntax error: `var` must bind the value to a **name**. `var x = 5`
  declares `x` and lets Go **infer** its type (`int`) from the initializer.

**Key detail:** with an initializer you don't write the type — `var x = 5` and
`var x int = 5` are equivalent, and idiomatic Go omits the redundant type.

**References**

- Go by Example — Variables: https://gobyexample.com/variables
