## concpat1 — worker pool

```go
func square(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for n := range jobs { // each worker pulls until jobs closes
        results <- n * n
    }
}
```

**Why it works**

- A fixed set of workers all `range` over the **same** `jobs` channel — Go hands
  each queued job to whichever worker is free, spreading the load. Closing `jobs`
  ends every worker's loop; `wg.Wait()` then knows they're done, so `results` can
  be closed and drained.

**Under the hood**

- There is no scheduler or work-stealing code here: workers parked on an empty
  channel sit in its receive queue, and each send wakes exactly one of them. That
  is why a slow job never idles the others — a worker that finishes early is
  already back in the queue asking for the next one.

**Common mistake**

- Closing `results` before `wg.Wait()`. A straggler still inside its loop then
  sends on a closed channel and the program panics. The order is fixed by
  ownership: `close(jobs)` ends the workers, `wg.Wait()` proves no worker can
  send again, and only then is `close(results)` safe.

**Key detail:** the worker count is a **limit**, not a speed setting. For CPU-bound
work `runtime.GOMAXPROCS(0)` is the sane default; for I/O-bound work N is the
concurrency cap on whatever you are calling — the reason a pool is the standard
way to avoid opening thousands of connections to a service that tolerates 50.

**See also:** concpat2 (fan-in's closer goroutine) · concpat3 (leak on early
exit) · channels2 (the directional parameters) · the
[patterns chapter](../README.md)

**References**

- Go blog — Pipelines and cancellation: https://go.dev/blog/pipelines
- pkg.go.dev — errgroup (pool + context + first error):
  https://pkg.go.dev/golang.org/x/sync/errgroup
- Go by Example — Worker Pools: https://gobyexample.com/worker-pools
