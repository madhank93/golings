# Deterministic Time (testing/synctest)

`testing/synctest` (GA in Go 1.25) runs test code in a "bubble" with a virtual
clock. When every goroutine in the bubble blocks, time jumps to the next timer —
so `time.Sleep` returns instantly and time-based tests become fast and
deterministic instead of slow and flaky.

- **synctest1** — wrap a test in `synctest.Test`; `synctest.Wait()` for idle
- **synctest2** — virtual time only advances when the test goroutine blocks too

## Resources

- [testing/synctest package](https://pkg.go.dev/testing/synctest)
- [Go blog: Testing concurrent code with synctest](https://go.dev/blog/synctest)
