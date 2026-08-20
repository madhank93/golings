# The maps package

The `maps` package (Go 1.21, extended in 1.23) is the map counterpart to
`slices`: generic helpers for whole-map operations that everyone used to
hand-roll — cloning, comparing, bulk deletion, and iterating keys and values.

Note the name collision. This chapter is about the **standard-library package**
`maps`; the `maps` *topic* earlier in the curriculum was about the built-in
`map[K]V` type. They are different things with the same word.

## 1. `maps.Clone`: because assignment does not copy

```go
cp := settings           // NOT a copy — both names, one map
cp := maps.Clone(settings)  // an independent shallow copy
```

A map value is a **pointer to a runtime hash table**. Assigning one copies the
pointer, so a write through either name is visible through the other — which is
`mapspkg1`, where a function that "copies then adds a default" silently mutates
its caller's map.

`Clone` is shallow: the new map has its own table, but values are copied as-is.
Clone a `map[string][]int` and both maps share the same slices, so appending
through one is still visible through the other. Deep copying is your job.

The same rule bites for **any function that takes a map and writes to it**. A
map parameter is a shared handle, not a snapshot — unlike a slice, where at
least an `append` cannot reach the caller.

## 2. `maps.DeleteFunc`: one pass instead of two

Deleting while ranging used to need a two-step, because the safe pattern was to
collect keys first:

```go
var expired []string
for k, v := range m {
    if v <= 0 { expired = append(expired, k) }
}
for _, k := range expired { delete(m, k) }
```

`mapspkg2` replaces all of it:

```go
maps.DeleteFunc(m, func(k string, v int) bool { return v <= 0 })
```

Worth knowing: deleting during a `range` was *already* safe in Go — an entry not
yet reached is simply not produced. The two-step existed mostly as folklore.
`DeleteFunc` is clearer regardless, and says the intent in one line.

## 3. Iterators: `Keys`, `Values`, `All`

Since Go 1.23 these return **iterators**, not slices:

```go
for k := range maps.Keys(m) { … }                // iter.Seq[K]
for k, v := range maps.All(m) { … }              // iter.Seq2[K, V]

keys := slices.Collect(maps.Keys(m))             // if you want a slice
sorted := slices.Sorted(maps.Keys(m))            // sorted, in one call
```

That last line is the answer to map iteration being randomised: `slices.Sorted(maps.Keys(m))`
is the idiomatic way to produce stable output. (`golang.org/x/exp/maps.Keys`,
which returned a slice directly, is the older API — do not confuse the two.)

## 4. The rest of the package

| Call | Does |
|---|---|
| `maps.Clone(m)` | shallow independent copy |
| `maps.Equal(a, b)` | same keys and `==` values |
| `maps.EqualFunc(a, b, eq)` | same, with a custom comparison |
| `maps.Copy(dst, src)` | merge `src` into `dst`, overwriting |
| `maps.DeleteFunc(m, del)` | delete matching entries |
| `maps.Keys` / `Values` / `All` / `Insert` / `Collect` | iterator bridges |

`maps.Equal` is worth calling out: maps are **not comparable** with `==` (only
against `nil`), so before this package every test comparing two maps used
`reflect.DeepEqual`. `maps.Equal` is type-safe, faster, and says what it means.

## Gotchas

- **`cp := m` shares the map.** Only `maps.Clone` copies it.
- **`Clone` is shallow** — nested slices, maps, and pointers stay shared.
- **`Clone(nil)` returns `nil`**, which is readable but panics on write.
- **`maps.Keys` returns an iterator, not a slice** (1.23+). Wrap with
  `slices.Collect` or `slices.Sorted`.
- **`maps.Copy(dst, src)` overwrites** existing keys in `dst`.
- **None of this is concurrency-safe.** A `Clone` racing a write is still a data
  race — see `safety2`.

## The exercises

- **mapspkg1** — replace `cp := settings` with `maps.Clone` so the caller's map
  is untouched.
- **mapspkg2** — collapse the collect-then-delete loop into `maps.DeleteFunc`.

## Source references

- [pkg.go.dev: maps](https://pkg.go.dev/maps) ·
  [slices](https://pkg.go.dev/slices)
- [Go 1.21 release notes: maps](https://go.dev/doc/go1.21#maps) ·
  [Go 1.23: iterators](https://go.dev/doc/go1.23#iterators)
- [Go blog: Go maps in action](https://go.dev/blog/maps)

**Next: [structured_logging](../structured_logging/) →** — the other package a
modern service reaches for on line one.
