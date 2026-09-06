## anonymous_functions1 — call the literal with an argument

```go
func(name string) {
    fmt.Printf("Hello %s", name)
}("Gopher")
```

**Why it works**

- A function literal declares a parameter, so invoking it needs a matching
  argument. The trailing `("Gopher")` is the call — without it the literal is
  just a value nobody uses, and the compiler says so.

**Common mistake**

- Dropping the parentheses entirely and expecting the body to run. A literal on
  its own is an expression of function type; Go has no bare-expression statement,
  so it fails with "evaluated but not used".

**Key detail:** invoking a literal on the spot is rare in Go. The forms you
actually write are assigning it, passing it (`sort.Slice`, `http.HandleFunc`),
returning it, or starting it — `go func(){…}()` and `defer func(){…}()` are the
same syntax with a keyword in front.

**See also:** anonymous_functions2 (matching a signature) ·
anonymous_functions3 (closures) · defer1 · the [chapter](../README.md)

**References**

- Go spec — Function literals: https://go.dev/ref/spec#Function_literals
- Effective Go — Function literals: https://go.dev/doc/effective_go#func_literals
