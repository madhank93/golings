## defer2 — cleanup on every path

```go
func process(r *Resource, early bool) {
    defer r.Close() // runs whichever way the function returns

    if early {
        return
    }
    // ... more work ...
}
```

**Why it works**

- A deferred call runs when the **function** returns, not when the block ends —
  so the early `return` and the normal one both release the resource. Without
  the `defer`, the early path leaks it.

**Under the hood**

- Deferred calls also run while a **panic** unwinds the stack, which is what makes
  `defer` the only reliable cleanup in Go: there is no `finally`, no destructor,
  no RAII. Since Go 1.14 most defers are open-coded (inlined at each return), so
  a straight-line `defer` costs a few nanoseconds.

**Common mistake**

- Deferring inside a loop. The stack unwinds at *function* return, not per
  iteration, so `for … { f, _ := os.Open(…); defer f.Close() }` holds every
  descriptor until the function ends. Move the body into its own function.

**Key detail:** the idiom is acquire → check the error → `defer` the release, in
that order. Deferring before the error check calls `Close` on a `nil` value. And
a deferred `Close` on a *writer* discards its error — capture it into a named
result when the write must be durable.

**See also:** defer1 (LIFO and argument evaluation) · sync1 (`defer mu.Unlock()`) ·
files1 · logingest6 (shutdown paths) · the [chapter](../README.md)

**References**

- Effective Go — Defer: https://go.dev/doc/effective_go#defer
- Go 1.14 release notes — open-coded defers: https://go.dev/doc/go1.14#runtime
