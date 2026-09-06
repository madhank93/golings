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

**Under the hood**

- The timer is registered with the runtime's timer heap when the `select`
  evaluates its case expressions, not when the case is taken. Before Go 1.23 that
  timer stayed alive until it fired — the old reason to prefer `time.NewTimer`
  plus `Stop` in a loop. Since 1.23 an unreachable timer is collectable
  immediately, so `time.After` in a loop no longer leaks; a reusable timer is now
  an optimisation, not a fix.

**Common mistake**

- Believing the timeout stopped the work. It did not: the producer goroutine is
  still running, and if it sends on an unbuffered channel nobody is receiving
  from any more, it parks forever — a leak. Give the producer a buffered channel
  or a context it watches.

**Key detail:** this pattern times out *one wait*. It cannot propagate across
function or process boundaries and it cannot tell the other side to stop. That is
`context`'s job — and `ctx.Done()` is just another case in the same `select`.

**See also:** select1 (choice) · context2 (`WithTimeout`, the propagating
version) · concpat3 (leaks from abandoned producers) · the
[select chapter](../README.md)

**References**

- Go 1.23 release notes — timer changes: https://go.dev/doc/go1.23#timer-changes
- pkg.go.dev — time.After: https://pkg.go.dev/time#After
- Go by Example — Timeouts: https://gobyexample.com/timeouts
