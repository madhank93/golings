## synctest2 — virtual time

```go
go func() {
    time.Sleep(time.Hour)
    close(done)
}()
<-done // block here, so the bubble advances its fake clock an hour
// time.Since(start) == time.Hour, instantly
```

**Why it works**

- Inside a synctest bubble, time is **virtual**: when *every* goroutine (including
  the test's) is durably blocked, the fake clock jumps to the next pending timer.
  The one-hour sleep then completes immediately — but only once the test
  goroutine itself blocks.

**Under the hood**

- The rule is precise: nothing can happen at the current instant if every bubbled
  goroutine is waiting on something only another bubbled goroutine could deliver,
  so the clock may safely skip ahead. `time.Now`, `time.Sleep`, `time.After`,
  timers, tickers and `context` deadlines all read this bubble clock, which
  starts at a fixed instant — so timestamps are reproducible too.

**Common mistake**

- Asserting on elapsed time without blocking, which is the bug as written here.
  The test goroutine is runnable, so the bubble is not fully blocked, the clock
  never advances, and `time.Since(start)` is 0. Busy CPU work has the same
  effect — the fake clock does not make computation instant.

**Key detail:** the elapsed time is **exactly** `time.Hour`, not approximately.
That is what makes cache expiry, retry backoff, rate limiters, heartbeats and
graceful-shutdown grace periods testable at all: microsecond runtime, zero
flake, and a schedule that is identical on every machine.

**See also:** synctest1 (`Wait` and the bubble) · context2 (deadlines you can now
test) · select2 (`time.After` under a real clock) · the
[synctest chapter](../README.md)

**References**

- pkg.go.dev — testing/synctest: https://pkg.go.dev/testing/synctest
- Go blog — Testing concurrent code with synctest: https://go.dev/blog/synctest
- Go issue #67434 — the proposal: https://github.com/golang/go/issues/67434
