## context2 — deadlines with WithTimeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
defer cancel()

select {
case <-time.After(d):
    return nil
case <-ctx.Done():
    return ctx.Err() // context.DeadlineExceeded
}
```

**Why it works**

- `context.WithTimeout` returns a context that cancels **itself** after the
  duration. When it fires, `ctx.Done()` closes and `ctx.Err()` reports
  `context.DeadlineExceeded`.

**Under the hood**

- `WithTimeout(parent, d)` is `WithDeadline(parent, time.Now().Add(d))`, and the
  resulting `timerCtx` is an ordinary `cancelCtx` with a `time.AfterFunc` that
  calls its own cancel. So a timeout and a manual cancel close the same channel
  in the same way — only `Err()` distinguishes them. A derived deadline can only
  shorten: `WithTimeout(ctx, time.Hour)` on a context with 5 s left still fires
  in 5 s.

**Common mistake**

- Comparing with `==`. By the time an error crosses a few layers it is usually
  wrapped, so use `errors.Is(err, context.DeadlineExceeded)`. (This exercise's
  test compares directly because nothing wraps it.)

**Key detail:** `DeadlineExceeded` means the timer fired; `Canceled` means someone
called `cancel`. Still `defer cancel()` on a `WithTimeout` — it releases the timer
and detaches the child immediately instead of at the deadline. Both errors
satisfy the `net.Error` timeout interface, which is why HTTP clients surface
"context deadline exceeded".

**See also:** select2 (a timeout that does *not* propagate) · context1
(cancellation mechanics) · httpsrv3 / logingest6 (deadlines in a server) · the
[context chapter](../README.md)

**References**

- pkg.go.dev — context.WithTimeout: https://pkg.go.dev/context#WithTimeout
- src/context/context.go — `timerCtx`:
  https://github.com/golang/go/blob/master/src/context/context.go
- Go blog — Context: https://go.dev/blog/context
