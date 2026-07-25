## if1 — return the larger with an if

```go
func bigger(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Why it works**

- When `a > b` the function returns `a` early; otherwise control falls through
  to `return b`. No `else` is needed — an early `return` handles the split.

**Nuance:** Go conditions take **no parentheses** but the braces are
**mandatory**, even for one line. There is no ternary (`?:`) operator — an `if`
is the idiomatic way to choose a value.

**References**

- Go by Example — If/Else: https://gobyexample.com/if-else
