## mapspkg1 — maps.Clone

```go
cp := maps.Clone(settings) // independent shallow copy
cp["theme"] = "dark"       // does NOT touch the caller's map
```

**Why it works**

- `cp := settings` does **not** copy a map — both names refer to the *same*
  underlying map, so a write through `cp` is visible through `settings`.
  `maps.Clone` (Go 1.21+) returns a genuinely separate map, so mutating `cp`
  leaves the original untouched.

**Nuance:** maps (like slices and pointers) are **reference types** — assigning
one shares it. `maps.Clone` is **shallow**: it copies the entries, but if the
values were themselves maps/slices/pointers, those are still shared. For "don't
leak my changes to the caller," clone at the boundary.

**References**

- pkg.go.dev — maps.Clone: https://pkg.go.dev/maps#Clone
