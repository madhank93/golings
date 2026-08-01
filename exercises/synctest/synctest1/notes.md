## synctest1 — run tests in a bubble

```go
synctest.Test(t, func(t *testing.T) {
    var count atomic.Int64
    for range 3 {
        go func() { count.Add(1) }()
    }
    synctest.Wait() // blocks until every goroutine in the bubble is idle
    // count is now reliably 3
})
```

**Why it works**

- `testing/synctest` (Go 1.25) runs the body in a **bubble**. `synctest.Wait()`
  blocks until every goroutine in the bubble is idle — a **deterministic**
  replacement for `time.Sleep`-and-hope when waiting on background goroutines.

**Key detail:** `Wait` only works **inside** a bubble; called outside `synctest.Test`
it panics. This kills flaky concurrency tests: instead of sleeping an arbitrary
duration and praying the goroutines finished, you wait for provable quiescence.

**References**

- pkg.go.dev — testing/synctest: https://pkg.go.dev/testing/synctest
- The Go Blog — Testing concurrent code with synctest: https://go.dev/blog/synctest
