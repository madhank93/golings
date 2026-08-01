## anonymous_functions2 — match a function-typed variable's signature

```go
var sayBye func(name string)
sayBye = func(name string) {
    fmt.Printf("Bye %s", name)
}
sayBye("Gopher")
```

**Why it works**

- `sayBye` is declared with type `func(name string)`. The value assigned to it
  must have the **exact same signature**, so the closure needs the `name string`
  parameter (and must use it).

**Key detail:** functions are **first-class values** in Go — you can store them in
variables, pass them as arguments, and return them. The variable's type is the
function signature, and assignment is type-checked against it.

**References**

- Go by Example — Closures: https://gobyexample.com/closures
