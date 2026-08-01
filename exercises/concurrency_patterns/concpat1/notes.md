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

**Key detail:** the ordering dance matters — close `jobs` after sending, `wg.Wait()`
for the workers, **then** close `results`. Closing `results` early would panic a
still-writing worker. A buffered `results` (size = len(inputs)) keeps workers from
blocking on a full channel.

**References**

- Go by Example — Worker Pools: https://gobyexample.com/worker-pools
