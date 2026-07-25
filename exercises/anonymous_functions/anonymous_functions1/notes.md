## anonymous_functions1 — call an anonymous function with an argument

```go
func(name string) {
    fmt.Printf("Hello %s", name)
}("Gopher")
```

**Why it works**

- An anonymous function is defined and **immediately called** by the trailing
  `("Gopher")`. That argument fills the `name` parameter.

**Nuance:** the `(...)` right after the closing brace is what *invokes* it — this
is an IIFE (immediately-invoked function expression). Without the call it's just
a function value sitting there, doing nothing.

**References**

- Go by Example — Closures: https://gobyexample.com/closures
