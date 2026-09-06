## mapspkg2 — maps.DeleteFunc

```go
func dropExpired(m map[string]int) {
    maps.DeleteFunc(m, func(k string, v int) bool { return v <= 0 })
}
```

**Why it works**

- `maps.DeleteFunc` walks the map and deletes every entry whose predicate is
  true — one pass, one line, replacing the collect-the-keys-then-delete loop.

**Under the hood**

- The two-step version existed largely as folklore: deleting **during** a `range`
  has always been safe in Go, since an entry not yet reached is simply not
  produced. `DeleteFunc` is clearer regardless, and states the intent instead of
  the mechanics.

**Common mistake**

- Collecting the keys and then forgetting to delete them, which is exactly what
  the broken version does — the slice is built and discarded, and the map is
  unchanged. Silent, and only a test notices.

**Key detail:** the function mutates the caller's map, which is the point here —
but it means a "filter" that should leave the input alone needs
`maps.Clone` first. The package also has `Equal` (maps are not comparable with
`==`), `Copy`, and the iterator bridges `Keys`/`Values`/`All`.

**See also:** mapspkg1 (`Clone`) · range2 (ranging maps) · maps3 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — maps.DeleteFunc: https://pkg.go.dev/maps#DeleteFunc
- Go 1.21 release notes: https://go.dev/doc/go1.21#maps
