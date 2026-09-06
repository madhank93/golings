## mapspkg1 — maps.Clone

```go
func withDefault(settings map[string]string) map[string]string {
    cp := maps.Clone(settings)
    cp["theme"] = "dark"
    return cp
}
```

**Why it works**

- `cp := settings` copies a **pointer to the same hash table**, so writing
  through `cp` mutates the caller's map. `maps.Clone` allocates a new table and
  copies the entries, giving a genuinely independent map.

**Under the hood**

- A map value is a pointer to a runtime structure, which is why a map parameter
  is a shared handle rather than a snapshot. Unlike a slice — where at least an
  `append` cannot reach the caller — every write to a map parameter is visible
  outside.

**Common mistake**

- Assuming `Clone` is deep. It is **shallow**: clone a `map[string][]int` and
  both maps point at the same slices, so appending through one is visible through
  the other. Nested copying is yours to do.

**Key detail:** `maps.Clone(nil)` returns `nil`, which reads fine and panics on
write. If the input may be nil and the result must be writable, `make` when the
clone comes back nil.

**See also:** mapspkg2 (`DeleteFunc`) · maps1 (map basics) · slices3 (the same
sharing, one level down) · the [chapter](../README.md)

**References**

- pkg.go.dev — maps.Clone: https://pkg.go.dev/maps#Clone
- Go blog — Go maps in action: https://go.dev/blog/maps
