## defer1 — deferred calls run LIFO

```go
func order() (seq []int) {
    defer func() { seq = append(seq, 1) }()
    defer func() { seq = append(seq, 2) }()
    defer func() { seq = append(seq, 3) }()
    return // seq becomes [3, 2, 1]
}
```

**Why it works**

- Deferred calls run when the function returns, in **last-in, first-out** order —
  so the `3` defer runs first, then `2`, then `1`, giving `[3, 2, 1]`.

**Key detail:** the deferred closures modify the **named return value** `seq`. A
`defer` can read and change named results *after* the `return` statement runs but
*before* the function actually returns — the mechanism behind `defer`-based error
wrapping. Note: a deferred call's **arguments** are evaluated immediately, but its
**body** runs at return.

**References**

- A Tour of Go — Defer: https://go.dev/tour/flowcontrol/12
- Go by Example — Defer: https://gobyexample.com/defer
