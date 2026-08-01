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

**Key detail:** distinguish the two errors — `DeadlineExceeded` (timeout) vs
`Canceled` (someone called `cancel`). Still `defer cancel()` even with a timeout,
to free the timer immediately if the work finishes early. A timeout is just a
deadline the runtime sets for you.

**References**

- pkg.go.dev — context: https://pkg.go.dev/context
- Go by Example — Context: https://gobyexample.com/context
