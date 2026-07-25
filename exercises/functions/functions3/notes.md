## functions3 — pass the argument the function expects

```go
func main() {
    call_me(10)
}

func call_me(num int) { ... }
```

**Why it works**

- `call_me` requires one `int`, but was called as `call_me()`. Passing `call_me(10)`
  supplies the missing argument.

**Nuance:** Go checks argument **count and type** at compile time — too few, too
many, or a wrong-typed argument all fail to build. There are no default
parameter values in Go; every declared parameter must be supplied.

**References**

- Go by Example — Functions: https://gobyexample.com/functions
