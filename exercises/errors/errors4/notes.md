## errors4 — recover a panic into an error

```go
func safeRun(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    fn()
    return nil
}
```

**Why it works**

- `recover()` only does something inside a **deferred** function; there it stops
  a panicking goroutine and returns the panic value. Assigning to the **named
  return** `err` turns the panic into an ordinary error the caller can handle.

**Nuance:** reserve panic/recover for truly exceptional cases (or a package
boundary that must not crash its caller) — **not** routine error handling, which
stays value-based. `recover` outside a deferred function returns `nil` and does
nothing.

**References**

- Go by Example — Recover: https://gobyexample.com/recover
- Effective Go — Recover: https://go.dev/doc/effective_go#recover
