## anonymous_functions2 — the literal must match the type

```go
var sayBye func(name string)

sayBye = func(name string) {
    fmt.Printf("Bye %s", name)
}
sayBye("Gopher")
```

**Why it works**

- `sayBye` is declared as `func(name string)`, so only a literal with exactly one
  `string` parameter can be assigned to it. The broken version assigns a `func()`
  and refers to an `n` that was never declared.

**Under the hood**

- Function types are **structural**: parameter types and result types, in order.
  The parameter *names* are documentation only — `func(name string)` and
  `func(s string)` are the same type — which is why the fix is about the
  signature's shape, not its wording.

**Common mistake**

- Assuming an unused variable is harmless. Go rejects declared-and-unused locals,
  so a `sayBye` that is assigned but never called still fails to build.

**Key detail:** this is what makes functions first-class in Go. Because the type
is the signature, a literal can be passed anywhere a `func(string)` is expected —
the basis of callbacks, middleware, and `http.HandlerFunc`.

**See also:** anonymous_functions1 · anonymous_functions3 (capturing state) ·
functions4 (signatures) · di3 · the [chapter](../README.md)

**References**

- Go spec — Function types: https://go.dev/ref/spec#Function_types
- Go by Example — Closures: https://gobyexample.com/closures
