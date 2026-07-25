## sync1 — guard shared state with a Mutex

```go
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```

**Why it works**

- `c.n++` is not atomic (read, add, write). With 1000 goroutines incrementing at
  once, some updates are lost. A `sync.Mutex` lets only **one** goroutine hold the
  lock at a time, so increments happen one-by-one and the total is exactly 1000.

**Nuance:** `defer c.mu.Unlock()` guarantees the lock is released on every return
path. Guard **both** reads and writes — `Value()` locks too, so it never observes
a half-updated state. Keep the mutex a **pointer receiver** concern; copying a
`Counter` (and its mutex) is a bug `go vet` warns about.

**References**

- Go by Example — Mutexes: https://gobyexample.com/mutexes
- sync package: https://pkg.go.dev/sync
