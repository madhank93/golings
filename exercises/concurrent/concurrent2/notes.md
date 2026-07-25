## concurrent2 — a mutex-guarded counter

```go
var counter int
var mu sync.Mutex
go func() {
    defer wg.Done()
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

**Why it works**

- 100 goroutines each `counter++`. Because `++` is read-modify-write (not atomic),
  the `sync.Mutex` serializes them so no increment is lost — the total is exactly
  100.

**Nuance:** run this exercise with `go test -race` — without the lock the detector
flags the concurrent writes as a data race. A mutex is the general answer; for a
single counter, `sync/atomic` (see sync3) is a lock-free alternative.

**References**

- Go by Example — Mutexes: https://gobyexample.com/mutexes
