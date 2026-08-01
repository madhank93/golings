## select2 — timeouts with time.After

```go
select {
case v := <-ch:
    return v, true
case <-time.After(d):
    return 0, false // d elapsed before a value arrived
}
```

**Why it works**

- `time.After(d)` returns a channel that delivers one value after `d`. Racing it
  against `<-ch` in a `select` means: take the value if it arrives first,
  otherwise give up when the timer fires.

**Key detail:** this is *the* Go timeout pattern. Caveat: `time.After` leaks its timer
until it fires, so in a hot loop prefer a reusable `time.NewTimer` (with `Stop`).
For cancellation that propagates across call boundaries, use `context` (next
topic) instead.

**References**

- Go by Example — Timeouts: https://gobyexample.com/timeouts
