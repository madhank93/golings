## functions1 — define the function you call

```go
func main() {
    call_me()
}

func call_me() {}
```

**Why it works**

- `main` calls `call_me()`, but it was never defined. Adding a `func call_me() {}`
  gives the call something to resolve to.

**Common mistake**

- Assuming the definition has to come *before* the call. It does not: the
  compiler reads every declaration in the package before checking any body, so
  order in the file is free and there are no forward declarations or headers.

**Key detail:** the name is the wrong shape for Go. `call_me` compiles, but Go
names are mixedCaps — `callMe` — and the **case of the first letter is access
control**: uppercase is exported from the package, lowercase is not. Linters
flag the underscore; the compiler enforces the case rule.

**See also:** functions2 (parameter types) · functions3 (arity) ·
the [chapter](../README.md)

**References**

- Go spec — Function declarations: https://go.dev/ref/spec#Function_declarations
- Effective Go — Names: https://go.dev/doc/effective_go#names
