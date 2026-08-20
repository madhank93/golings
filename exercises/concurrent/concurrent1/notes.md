## concurrent1 — goroutines + WaitGroup

```go
var wg sync.WaitGroup
var mu sync.Mutex
for i := 0; i < goroutines; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        mu.Lock()
        fmt.Fprintf(buf, "Hello from goroutine %d!\n", i)
        mu.Unlock()
    }(i)
}
wg.Wait()
```

**Why it works**

- `go func(){...}()` starts a **goroutine** (a lightweight concurrent task). A
  `sync.WaitGroup` counts them: `Add(1)` before each, `Done()` when it finishes,
  and `Wait()` blocks until all are done. The `mu.Lock()` serializes the shared
  `buf` so the writes don't race.

**Under the hood**

- The `WaitGroup` is a counter plus a semaphore. `Add`/`Done` are atomic adds on
  one word; `Wait` parks the calling goroutine, and the `Done` that takes the
  counter to zero wakes it. Nothing spins. Each goroutine costs a 2 KB stack
  that grows by copying, which is why thousands of them are unremarkable.

**Common mistake**

- Calling `wg.Add(1)` *inside* the goroutine. It races with `Wait`: the counter
  may still be 0 when `Wait` runs, so `Wait` returns while the goroutines are
  still writing. `Add` before `go`, always.

**Key detail:** `defer wg.Done()` on the first line so early returns and panics
still decrement — a missed `Done` deadlocks `Wait` forever. Passing `i` as an
argument was mandatory before Go 1.22, when all iterations shared one loop
variable; since 1.22 each iteration gets its own, so the plain closure is also
correct. Output order is never guaranteed — assert on content, not sequence.

**See also:** concurrent2 (the race without the lock) · sync1 (mutex in a
struct) · safety1 (`wg.Go`, Go 1.25) · the [goroutines chapter](../README.md)

**References**

- The Go Memory Model: https://go.dev/ref/mem
- pkg.go.dev — sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
- A Tour of Go — Goroutines: https://go.dev/tour/concurrency/1
