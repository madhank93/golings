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

**Nuance:** pass the loop variable as an argument — `go func(i int){...}(i)` —
so each goroutine captures its **own** `i`. Also: without the `WaitGroup`, `main`
would return before the goroutines run; without the mutex, concurrent writes to
`buf` are a data race (`go test -race` would flag it).

**References**

- A Tour of Go — Goroutines: https://go.dev/tour/concurrency/1
- Go by Example — Goroutines: https://gobyexample.com/goroutines
