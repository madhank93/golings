## applied2 — a concurrency-safe store

```go
type Store struct {
    mu sync.Mutex
    m  map[string]int
}
func (s *Store) Set(key string, v int) { s.mu.Lock(); defer s.mu.Unlock(); s.m[key] = v }
func (s *Store) Get(key string) (int, error) {
    s.mu.Lock(); defer s.mu.Unlock()
    v, ok := s.m[key]
    if !ok { return 0, ErrNotFound }
    return v, nil
}
```

**Why it works**

- This pulls together the fundamentals: a **map** for storage, a **mutex** so
  concurrent `Set`/`Get` don't race, **methods** for a clean API, and a
  **sentinel error** (`ErrNotFound`) for the missing-key case that callers match
  with `errors.Is`.

**Nuance:** maps are **not** safe for concurrent use — a bare concurrent map
write can crash the program, so every access goes through the mutex. Returning
`(value, error)` instead of the map's comma-ok turns "missing" into a first-class
error the caller handles like any other.

**References**

- The Go Blog — Go maps in action: https://go.dev/blog/maps
- pkg.go.dev — sync.Mutex: https://pkg.go.dev/sync#Mutex
