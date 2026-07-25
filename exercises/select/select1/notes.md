## select1 — wait on multiple channels

```go
select {
case v := <-a:
    return v
case v := <-b:
    return v
}
```

**Why it works**

- `select` blocks until **one** of its cases can proceed, then runs that case. So
  `firstReady` returns whichever of `a` or `b` has a value first.

**Nuance:** if several cases are ready at once, `select` picks one **at random** —
never top-to-bottom — so you can't rely on ordering for fairness. A `select` with
no ready case and no `default` blocks; an **empty** `select {}` blocks forever.

**References**

- Go by Example — Select: https://gobyexample.com/select
