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
  the test's) is blocked, the fake clock jumps to the next pending timer. So a
  one-hour `time.Sleep` completes instantly — but only once the test goroutine
  itself blocks (here, on `<-done`).

**Nuance:** the subtlety is that virtual time advances only when the whole bubble
is idle. Read the clock **without** blocking and no time passes (elapsed = 0).
This makes timeout/retry/ticker tests exact and instant instead of slow and
flaky.

**References**

- The Go Blog — Testing concurrent code with synctest: https://go.dev/blog/synctest
