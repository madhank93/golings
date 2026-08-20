## variables1 — a variable needs a name

```go
var x = 5
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `var = 5` is a syntax error: `var` must bind the value to a **name**. `var x = 5`
  declares `x` and lets Go **infer** its type (`int`) from the initializer.

**Common mistake**

- Writing the type as well out of habit — `var x int = 5`. It compiles, but the
  type is redundant next to a literal that already implies it, and `gofmt`'s
  companion linters will point at it.

**Key detail:** there are four declaration forms and they build the same variable:
`var x int` (zero value), `var x int = 5`, `var x = 5` (type inferred), and
`x := 5` (short form, functions only). Inference reads the *initializer*, so
`var x = 5` is an `int` and `var x = 5.0` is a `float64`.

**See also:** variables2 (`:=`) · variables3 (no type, no value) ·
primitive_types5 (the real type names) · the [chapter](../README.md)

**References**

- Go spec — Variable declarations: https://go.dev/ref/spec#Variable_declarations
- Go by Example — Variables: https://gobyexample.com/variables
