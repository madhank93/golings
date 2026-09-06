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

- `testing/synctest` (GA in Go 1.25) runs the body in a **bubble**.
  `synctest.Wait()` returns once every *other* goroutine in the bubble is
  durably blocked or finished, so the assertion runs at a known point in the
  schedule instead of a hoped-for one.

**Under the hood**

- The bubble tracks the goroutines started inside it and knows, for each, whether
  it is blocked on something only another bubbled goroutine could satisfy —
  *durably blocked*. Channel operations, `WaitGroup.Wait`, mutexes and
  `time.Sleep` qualify; I/O and channels shared with the outside world do not,
  because the runtime cannot know when they will complete.

**Common mistake**

- Calling `synctest.Wait()` outside a bubble — it panics, which is exactly how
  this exercise starts. The bubble is a boundary: everything that participates
  must be started inside `synctest.Test`.

**Key detail:** this replaces `time.Sleep(50*time.Millisecond)`-and-hope. That
sleep is a guess that is simultaneously too short on loaded CI (flake) and too
long everywhere else (slow suite). `Wait` is deterministic and instant, and
`-race` still works normally inside the bubble.

**See also:** synctest2 (the fake clock) · concurrent1 (`WaitGroup` as the
non-test alternative) · testadv1 (subtests) · the
[synctest chapter](../README.md)

**References**

- pkg.go.dev — testing/synctest: https://pkg.go.dev/testing/synctest
- Go blog — Testing concurrent code with synctest: https://go.dev/blog/synctest
- Go 1.25 release notes: https://go.dev/doc/go1.25#testing-synctest
