## functions2 — parameters need types

```go
func callMe(num int) {
    for n := 0; n <= num; n++ { ... }
}
```

**Why it works**

- `func callMe(num)` is invalid — every parameter must declare its type. `num int`
  fixes it.

**Nuance:** when consecutive parameters share a type you may write it once —
`func add(a, b int)` means both `a` and `b` are `int`. But the type can never be
omitted entirely.

**References**

- Go by Example — Functions: https://gobyexample.com/functions
