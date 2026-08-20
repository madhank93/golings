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

**Under the hood**

- `counter++` compiles to load, add, store. Two goroutines that load the same
  value both store the same result and one increment disappears — no crash, just
  a number that is quietly too small. `Lock` is a single compare-and-swap when
  uncontended; only a loser parks on the mutex's semaphore.

**Common mistake**

- Assuming the test passing proves the code is correct. Without the lock this
  test *sometimes* prints 100 on a fast machine. The race is real either way —
  `go test -race` reports it deterministically, which is why `mise run test`
  always passes `-race`.

**Key detail:** a mutex is the general answer, because it can span several fields
at once. For this single word, `atomic.Int64.Add(1)` (sync3) does the same job
lock-free. The moment two fields must stay consistent with each other, no atomic
can help and you need the mutex.

**See also:** sync1 (the same lock as a struct field) · sync3 (atomic
alternative) · concurrent3 (avoid sharing entirely) · the
[goroutines chapter](../README.md)

**References**

- Data Race Detector: https://go.dev/doc/articles/race_detector
- The Go Memory Model: https://go.dev/ref/mem
- Go by Example — Mutexes: https://gobyexample.com/mutexes
