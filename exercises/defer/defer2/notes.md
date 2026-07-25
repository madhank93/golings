## defer2 — defer guarantees cleanup

```go
func process(r *Resource, early bool) {
    defer r.Close() // runs no matter which return fires
    if early {
        return
    }
}
```

**Why it works**

- `defer r.Close()` is scheduled the moment it's reached, so it runs on **every**
  return path — the early one and the normal one. Both `r1` and `r2` end closed.

**Nuance:** this is *the* Go idiom — put the cleanup (`Close`, `Unlock`, `Done`)
right next to the acquisition, and `defer` makes it fire even on panics or
multiple return points. It keeps open/close paired and unmissable.

**References**

- Go by Example — Defer: https://gobyexample.com/defer
- Effective Go — Defer: https://go.dev/doc/effective_go#defer
