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

- One copier goroutine per input channel forwards values into a shared `out`.
  Each copier knows when *its* input ended; the `WaitGroup` is what tells the
  closer that **all** of them did.

**Under the hood**

- The closer has to be its own goroutine: `merge` must return `out` immediately so
  the caller can start receiving, and `wg.Wait()` cannot complete until the caller
  drains the copiers' sends. Blocking on `Wait` inside `merge` would deadlock —
  the caller is waiting for the return, the copiers are waiting for the caller.

**Common mistake**

- Expecting input order in the output. Several copiers race to send on `out`, so
  the interleaving is arbitrary and varies per run. If order matters, tag values
  with an index and sort after collecting — fan-in is the wrong shape for
  ordered work.

**Key detail:** the return type is `<-chan int`, receive-only. The function that
creates a channel owns it, and handing back the restricted type makes it a
compile error for the caller to close it — the ownership rule enforced by the
type system rather than by a comment.

**See also:** concpat1 (pool) · concpat3 (pipeline) · safety3 (why exactly one
closer) · channels2 (directional types) · the [patterns chapter](../README.md)

**References**

- Go blog — Pipelines and cancellation: https://go.dev/blog/pipelines
- Go spec — Channel types: https://go.dev/ref/spec#Channel_types
- pkg.go.dev — sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
