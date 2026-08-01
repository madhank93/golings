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

- A plain `map` crashes on concurrent writes ("fatal error: concurrent map
  writes"). `sync.Map`'s `Store`/`Load` are safe for simultaneous use, so many
  goroutines can populate the registry at once.

**Key detail:** `sync.Map` stores `any`, so `Load` returns `(any, bool)` and you
type-assert `v.(int)`. Reach for it only in its niche — **stable keys with lots
of concurrent reads** (caches, registries). For most code a `map` + `RWMutex`
(safety1) is clearer and just as fast.

**References**

- pkg.go.dev — sync.Map: https://pkg.go.dev/sync#Map
