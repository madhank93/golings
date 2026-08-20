# Deterministic time (testing/synctest)

Concurrent code is testable right up to the moment time gets involved. Then the
tests grow `time.Sleep(50 * time.Millisecond)` — long enough that the goroutine
"should" be done, short enough that the suite still finishes. Both halves of that
sentence are guesses. On a loaded CI box the sleep is too short and the test
flakes; on your laptop the suite spends its life asleep. Testing a one-hour
retry backoff this way is simply not possible.

`testing/synctest` (experimental in Go 1.24, GA in **Go 1.25**) removes the
guess. It runs a test inside a *bubble* with a **fake clock**: when every
goroutine in the bubble is blocked, there is nothing left to wait for, so the
clock jumps straight to the next timer. A one-hour sleep returns instantly, and
"instantly" is exact rather than probable.

## 1. The bubble

```go
func TestThing(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        // everything in here, and every goroutine started from here,
        // is inside the bubble
    })
}
```

`synctest.Test` starts the bubble, runs the function as its root goroutine, and
returns when that function and every goroutine it started have finished. Inside,
`time.Now`, `time.Sleep`, `time.After`, `time.NewTimer`, `time.Tick` and
`context` deadlines all read the bubble's fake clock instead of the wall clock.
The clock starts at a fixed instant, so timestamps are reproducible too.

The bubble is a boundary, not a mode switch: goroutines started inside are in it,
the world outside is not. A channel shared with a goroutine outside the bubble is
a mistake — the runtime cannot tell whether that goroutine will eventually send,
so the bubble's central question becomes unanswerable and it panics rather than
guess.

## 2. Durably blocked — the one concept

Time advances only when **every** goroutine in the bubble is *durably blocked*:
blocked on something that only another bubbled goroutine could unblock.

Durably blocking:

- channel send/receive where the other end is in the bubble
- `select` where every case is such an operation
- `sync.WaitGroup.Wait`, `sync.Mutex.Lock`, `sync.Cond.Wait`
- `time.Sleep`

Not durably blocking:

- blocking on I/O — a network read, a file, a syscall
- waiting on a channel shared with a goroutine outside the bubble
- spinning on the CPU, or `runtime.Gosched()`

The distinction is exactly the one the fake clock needs. If the only way forward
is another bubbled goroutine, and all of them are stopped, then nothing will ever
happen at the current time — so the clock may safely skip to the next timer. If
something outside might still deliver a byte, it may not.

`synctest2` is that rule as a failing test. A goroutine sleeps an hour and closes
a channel; the test then measures elapsed time *without waiting for it*. The test
goroutine is running, so the bubble is not fully blocked, so the clock never
moves and `time.Since(start)` is 0. Add `<-done` and both goroutines are durably
blocked, the clock jumps an hour, the sleeper wakes, and the elapsed time is
**exactly** `time.Hour` — not "about an hour", exactly.

## 3. `synctest.Wait`

```go
for range 3 {
    go func() { count.Add(1) }()
}
synctest.Wait()          // returns when the other bubbled goroutines are blocked or done
```

`Wait` blocks the calling goroutine until every *other* goroutine in the bubble
is durably blocked or finished. It is the deterministic replacement for "sleep a
bit and hope the workers got there" — the assertion after it runs at a known
point in the schedule, every run, on every machine.

It only works inside a bubble; calling it from an ordinary test panics, which is
what `synctest1` starts out doing.

## 4. What this is for

The pattern generalises to everything time-shaped that used to be untestable or
slow: cache expiry, retry with exponential backoff, rate limiters, heartbeats,
`context.WithTimeout` paths, graceful-shutdown grace periods.

```go
synctest.Test(t, func(t *testing.T) {
    c := NewCache(time.Hour)
    c.Set("k", "v")
    time.Sleep(2 * time.Hour)   // instant
    if _, ok := c.Get("k"); ok {
        t.Error("entry should have expired")
    }
})
```

The result runs in microseconds and cannot flake, because there is no race
between the timer and the assertion — the schedule is deterministic. Combine it
with `-race`, which still works normally inside a bubble.

## Gotchas

- **The bubble panics on I/O it cannot account for**, including a `net/http`
  call to a real server. Use `httptest` or a fake inside the bubble.
- **`synctest.Wait` outside a bubble panics.** So does sharing a channel with an
  unbubbled goroutine.
- **The fake clock does not make CPU work instant.** A busy loop is not blocked,
  so the clock stands still while it spins.
- **The bubble's clock is per-bubble**; two bubbles do not share time, and
  neither reflects the wall clock.
- **The API changed between 1.24 and 1.25.** The experimental version was
  `synctest.Run(func())` behind `GOEXPERIMENT=synctest`; the GA version is
  `synctest.Test(t, func(*testing.T))`. This repo is on Go 1.26 — use `Test`.

## The exercises

- **synctest1** — wrap a test body in `synctest.Test` so `synctest.Wait()` has a
  bubble to work in, and assert on three goroutines without sleeping.
- **synctest2** — make the test goroutine block so the fake clock advances, and
  watch a one-hour sleep resolve to exactly one hour, instantly.

## Source references

- [pkg.go.dev: testing/synctest](https://pkg.go.dev/testing/synctest) — the
  precise definition of "durably blocked"
- [Go blog: Testing concurrent code with synctest](https://go.dev/blog/synctest)
- [Go 1.25 release notes](https://go.dev/doc/go1.25#testing-synctest) — GA, and
  the `Run` → `Test` change
- [Go issue #67434](https://github.com/golang/go/issues/67434) — the proposal
  and its design discussion

**End of the concurrency series.** Next in the curriculum:
[stdlib_essentials](../stdlib_essentials/) — and the
[capstone](../capstone/) builds a concurrent log-ingest service out of every
chapter here.
