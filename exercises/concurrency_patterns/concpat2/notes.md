## concpat2 — fan-in

```go
func merge(chans ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range chans {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c { out <- v }
        }(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```

**Why it works**

- One goroutine per input channel copies its values into the shared `out`. A
  separate goroutine waits for **all** of them (`wg.Wait()`) and then closes
  `out` — so the consumer's `range` ends exactly once every input is drained.

**Nuance:** the closer must run in its **own** goroutine — calling `wg.Wait()`
inline would block `merge` from returning. Pass `c` as an argument to the
goroutine so each captures its own channel, not the loop variable.

**References**

- The Go Blog — Pipelines and cancellation: https://go.dev/blog/pipelines
