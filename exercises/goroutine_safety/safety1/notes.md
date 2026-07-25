## safety1 — RWMutex for read-heavy state

```go
func (c *Counter) Inc()   { c.mu.Lock();  defer c.mu.Unlock();  c.n++ }
func (c *Counter) Value() int { c.mu.RLock(); defer c.mu.RUnlock(); return c.n }
```

**Why it works**

- `sync.RWMutex` has two locks: `Lock` (writer, exclusive) and `RLock` (reader,
  shared). Many readers can hold `RLock` at once, but a writer's `Lock` waits for
  all readers to finish. `Inc` writes, so it takes the write lock; `Value` reads,
  so it takes the cheaper read lock.

**Nuance:** `RWMutex` pays off only when reads **greatly** outnumber writes —
otherwise a plain `Mutex` is simpler and often faster (RWMutex has more
bookkeeping). Either way, every access must be guarded; run with `-race` to prove
it.

**References**

- pkg.go.dev — sync.RWMutex: https://pkg.go.dev/sync#RWMutex
- Data Race Detector: https://go.dev/doc/articles/race_detector
