## functions2 — parameters declare their type

```go
func callMe(num int) {
    for n := 0; n <= num; n++ {
        fmt.Printf("Num is %d\n", n)
    }
}
```

**Why it works**

- `func callMe(num)` names a parameter without a type. Go infers types for
  *variables* from their initializer, but a parameter has no initializer — the
  signature is the only place its type can come from.

**Common mistake**

- Expecting inference from the call site. `callMe(10)` looks like enough
  information, but a signature is a contract compiled once, independent of who
  calls it — otherwise every caller could redefine the function's meaning.

**Key detail:** parameters sharing a type may share the annotation:
`func add(a, b int) int` is the same as `func add(a int, b int) int`. And the
type comes *after* the name — the order that makes Go's declarations read left
to right instead of spiralling the way C's do.

**See also:** functions3 (arity) · functions4 (return types) ·
the [chapter](../README.md)

**References**

- Go spec — Function types: https://go.dev/ref/spec#Function_types
- Go by Example — Functions: https://gobyexample.com/functions
