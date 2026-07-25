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

**Nuance:** Go has **no forward declarations** and order doesn't matter — a
function can be defined anywhere in the package (above or below its callers) as
long as it exists somewhere.

**References**

- Go by Example — Functions: https://gobyexample.com/functions
