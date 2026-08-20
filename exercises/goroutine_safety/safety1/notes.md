## safety1 — RWMutex for read-heavy state

```go
func (c *Counter) Inc()      { c.mu.Lock();  defer c.mu.Unlock();  c.n++ }
func (c *Counter) Value() int { c.mu.RLock(); defer c.mu.RUnlock(); return c.n }
```

**Why it works**

- `sync.RWMutex` has two locks: `Lock` (writer, exclusive) and `RLock` (reader,
  shared). Many readers can hold `RLock` at once, but a writer's `Lock` waits for
  all readers to finish. `Inc` writes, so it takes the write lock; `Value` reads,
  so it takes the cheaper read lock.

**Under the hood**

- The lock keeps a reader count and two semaphores. A waiting writer immediately
  blocks *new* readers by driving the reader count negative, then waits only for
  the readers already inside. That write-preference is what stops a steady stream
  of readers from starving a writer forever.

**Common mistake**

- Taking `RLock` recursively. If a reader holds `RLock`, a writer queues behind
  it, and that reader takes `RLock` again, the second acquisition waits behind
  the writer, which waits behind the first — a deadlock written as two harmless
  reads. Neither `RLock` nor `Lock` is reentrant.

**Key detail:** `RWMutex` pays off only when reads **greatly** outnumber writes and
the critical section is long enough to matter — it tracks more state than a plain
`Mutex`, so for short sections a `Mutex` is often faster. Measure with
`go test -bench` instead of assuming. This exercise also uses `wg.Go` (Go 1.25),
which folds `Add(1)`, `go`, and `defer Done()` into one call.

**See also:** sync1 (plain mutex) · safety2 (`sync.Map` as the other answer) ·
concurrent2 (the underlying race) · the [safety chapter](../README.md)

**References**

- src/sync/rwmutex.go: https://github.com/golang/go/blob/master/src/sync/rwmutex.go
- pkg.go.dev — sync.RWMutex: https://pkg.go.dev/sync#RWMutex ·
  WaitGroup.Go: https://pkg.go.dev/sync#WaitGroup.Go
- Data Race Detector: https://go.dev/doc/articles/race_detector
