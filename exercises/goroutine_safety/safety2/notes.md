## safety2 — sync.Map

```go
func (r *Registry) Set(name string, id int) { r.m.Store(name, id) }
func (r *Registry) Get(name string) (int, bool) {
    v, ok := r.m.Load(name)
    if !ok { return 0, false }
    return v.(int), true
}
```

**Why it works**

- A plain `map` is not safe for concurrent writes. `sync.Map`'s `Store`/`Load`
  are safe for simultaneous use, so many goroutines can populate the registry at
  once.

**Under the hood**

- Concurrent writes to a plain map do not merely corrupt it — the runtime detects
  the situation and aborts the process with `fatal error: concurrent map writes`.
  That is not a panic and `recover` cannot catch it: a map write spans bucket,
  count and growth state, and the runtime refuses to continue from a half-updated
  one. `sync.Map` avoids the abort with an implementation built for concurrency —
  historically a lock-free read map with a mutex-guarded dirty map behind it, and
  since Go 1.24 a concurrent hash-trie.

**Common mistake**

- Reaching for `sync.Map` as the default concurrent map. It stores `any`, so you
  lose type safety, pay a boxing allocation, get no `len`, and `Range` sees only
  a snapshot. A `map[K]V` behind an `RWMutex` (safety1) is usually clearer and
  faster.

**Key detail:** use it in the two cases its own docs name — a key written once and
read many times, or disjoint goroutines touching disjoint keys (caches,
registries). `Load` returns `(any, bool)`, so a wrong type assertion panics at
the *read* site, far from the `Store` that caused it.

**See also:** safety1 (`RWMutex` + map) · sync1 (the lock it replaces) ·
maps1 (comma-ok on a plain map) · the [safety chapter](../README.md)

**References**

- pkg.go.dev — sync.Map: https://pkg.go.dev/sync#Map
- Go 1.24 release notes — the hash-trie rewrite: https://go.dev/doc/go1.24#sync
- src/sync/map.go: https://github.com/golang/go/blob/master/src/sync/map.go
