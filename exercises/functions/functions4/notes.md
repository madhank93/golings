## functions4 — declare the return type

```go
func addNumbers(a int, b int) int {
    return a + b
}
```

**Why it works**

- The function returns `a + b`, so its signature must declare that it returns an
  `int`. Without the `int` after the parameter list, `return a + b` is a compile
  error.

**Key detail:** the return type goes **after** the parameter list. Go functions can
return **multiple** values too — `func vals() (int, int)` — which is how Go
idiomatically returns a result alongside an `error`.

**References**

- Go by Example — Functions: https://gobyexample.com/functions
- Go by Example — Multiple Return Values: https://gobyexample.com/multiple-return-values
