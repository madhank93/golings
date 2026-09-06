## applied2 — a concurrency-safe store

```go
func (s *Store) Set(key string, val int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[key] = val
}
```

**Why it works**

- The mutex serialises every access to the map, so concurrent `Set` and `Get`
  calls cannot race. `Get` uses comma-ok and returns `ErrNotFound` for a missing
  key, so callers can tell "absent" from "stored zero".

**Under the hood**

- Six ideas in twenty lines: the mutex sits next to the data it guards; the
  receivers are **pointers**, because a value receiver would copy the lock and
  guard nothing; `NewStore` exists because a `nil` map panics on write; the map
  is **unexported**, so there is no way in that skips the lock; comma-ok
  distinguishes absent from zero; and a sentinel error lets callers use
  `errors.Is`.

**Common mistake**

- Returning the internal map (or a slice it holds) from a getter. That hands out
  the very thing the lock protects — the caller then reads it unsynchronised
  while another goroutine writes. Return a copy.

**Key detail:** this is the shape of every cache, registry, and in-memory index in
Go. Scaling it up: `RWMutex` when reads dominate (safety1), sharding when one
lock becomes the bottleneck, and `context` on the methods once anything can
block. Run it with `-race` — an unlocked path is invisible until it is not.

**See also:** sync1 (mutex) · safety1 (`RWMutex`) · maps3 (comma-ok) ·
errors2 (sentinels) · logingest2 · the [chapter](../README.md)

**References**

- pkg.go.dev — sync.Mutex: https://pkg.go.dev/sync#Mutex
- Data Race Detector: https://go.dev/doc/articles/race_detector
