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

**Nuance:** `default` turns a blocking channel op into a **poll**. Handy for "take
it if it's there," but beware busy-looping — a `for { select { ... default: } }`
with no pause spins the CPU. Use it for a single try, not a tight loop.

**References**

- Go by Example — Non-Blocking Channel Operations: https://gobyexample.com/non-blocking-channel-operations
