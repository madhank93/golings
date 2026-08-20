## select3 — non-blocking with default

```go
select {
case v := <-ch:
    return v, true
default:
    return 0, false // nothing ready — don't block
}
```

**Why it works**

- A `default` case runs immediately when **no** other case is ready, so `select`
  never blocks. `tryReceive` reports `(value, true)` if a value is waiting, else
  `(0, false)`.

**Under the hood**

- `default` makes `selectgo` run only its first pass — poll the cases in random
  order, and if none can proceed, take the default instead of enqueuing and
  parking. That is why it is cheap for a single attempt and disastrous in a tight
  loop: each iteration is a full poll, and the goroutine never yields the CPU.

**Common mistake**

- `for { select { case v := <-ch: …; default: } }` — a spin loop that pegs a core
  asking "anything yet?" millions of times a second. If you want to wait, delete
  the `default` and let the runtime park the goroutine.

**Key detail:** `default` also works on the send side — the standard way to drop
work rather than block when a buffer is full (metrics, best-effort
notifications). Either way the answer is a **sample**: it is stale the instant it
returns, so never build logic on repeated polling.

**See also:** select1 (blocking choice) · channels1 (what "full" means) ·
safety1 (state you poll usually wants a lock instead) · the
[select chapter](../README.md)

**References**

- Go spec — Select statements: https://go.dev/ref/spec#Select_statements
- Go by Example — Non-Blocking Channel Operations: https://gobyexample.com/non-blocking-channel-operations
