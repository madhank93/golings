## sync1 — guard state with sync.Mutex

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```

**Why it works**

- `Lock` blocks until no other goroutine holds the mutex, so only one goroutine
  at a time runs `c.n++`. 1000 concurrent `Inc` calls therefore total exactly
  1000, and `-race` is silent.

**Under the hood**

- A `Mutex` is a state word plus a semaphore. Uncontended `Lock` is a single
  compare-and-swap that never enters the kernel; a loser spins briefly (the
  holder is often about to release) and then parks. If a waiter has been queued
  for more than 1 ms the mutex switches to *starvation mode* and hands ownership
  directly to the head of the queue, so no goroutine is starved indefinitely.

**Common mistake**

- Calling a locking method from another locking method. `sync.Mutex` is **not
  reentrant** — the second `Lock` deadlocks against your own goroutine. Factor
  the shared part into an unexported helper that assumes the lock is already
  held.

**Key detail:** a mutex protects **data, not code**, so keep it adjacent to the
field it guards and name that field in a comment. `defer Unlock` immediately
after `Lock` covers every return path and panic. Never copy a struct holding a
mutex — the copy has its own lock; `go vet` catches most of these.

**See also:** concurrent2 (the same fix, free-standing) · sync3 (atomic when the
state is one word) · safety1 (`RWMutex` for read-heavy state) · the
[sync chapter](../README.md)

**References**

- src/sync/mutex.go — state bits, spinning, starvation mode:
  https://github.com/golang/go/blob/master/src/sync/mutex.go
- The Go Memory Model — locks: https://go.dev/ref/mem#locks
- pkg.go.dev — sync.Mutex: https://pkg.go.dev/sync#Mutex
