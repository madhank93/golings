## variables2 — declare before you assign

```go
x := 5
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `x = 5` assigns to an `x` that was never declared → compile error. `:=`
  **declares and assigns** in one step, so `x` now exists.

**Nuance:** `:=` works **only inside a function**. At package level, or when a
name already exists, use `var x = 5` (declare) or plain `x = 5` (reassign an
existing variable). `:=` always introduces at least one new variable.

**References**

- Go by Example — Variables: https://gobyexample.com/variables
