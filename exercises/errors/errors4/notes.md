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

- A panic unwinds the stack running deferred calls. `recover()` inside one of
  them stops the unwinding and returns the panic value; assigning to the **named
  result** `err` makes `safeRun` return an error instead of dying.

**Under the hood**

- Three requirements, all load-bearing: `recover()` must be called **directly by
  a deferred function** (nested one level deeper it returns `nil`); the result
  must be **named**, or the closure has nothing to assign to; and `return x`
  assigns the result *before* defers run, which is what lets the closure
  overwrite it.

**Common mistake**

- Using panic/recover as control flow. It is for **programmer error** — broken
  invariants, impossible states — not for a missing file or a bad request.
  Returning an error is cheaper, visible in the signature, and checkable.

**Key detail:** the panic value is an `any`, so format it (`%v`) rather than
treating it as an error. And a panic in **any** goroutine kills the process — a
`recover` in `main` cannot save you from one in a worker, so each goroutine that
needs protection needs its own.

**See also:** defer1 (named results) · defer2 · errors1 · safety3 (panics from
misused channels) · the [chapter](../README.md)

**References**

- Go blog — Defer, Panic, and Recover: https://go.dev/blog/defer-panic-and-recover
- Go spec — Handling panics: https://go.dev/ref/spec#Handling_panics
