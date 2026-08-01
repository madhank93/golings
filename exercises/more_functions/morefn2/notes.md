## morefn2 — variadic parameters

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

**Why it works**

- `nums ...int` lets `sum` accept any number of ints. Inside the function `nums`
  is just an `[]int`, so you `range` over it.

**Key detail:** call it with individual args (`sum(1, 2, 3)`) **or** spread an
existing slice with `...` (`sum(nums...)`). The variadic parameter must be the
**last** one, and calling with none yields an empty (not nil-panicking) slice.

**References**

- Go by Example — Variadic Functions: https://gobyexample.com/variadic-functions
