# Maps

A map is Go's hash table: `map[K]V` associates keys of one type with values of
another, with average O(1) lookup, insert, and delete. Two of its behaviours
catch newcomers out, and both are deliberate — reading a missing key returns the
**zero value** rather than an error, and iteration order is **deliberately
randomised**.

These three exercises cover declaring one, filling one from a literal, and the
missing-key case that quietly returns `""`.

## 1. Creating and using

```go
m := make(map[string]int)          // empty, ready to write
ages := map[string]int{            // literal
    "John": 30,
    "Ana":  21,
}

ages["Carl"] = 42      // insert or overwrite
n := ages["John"]      // read: 30
delete(ages, "Ana")    // remove; no-op if absent
len(ages)              // 2
```

Both type parameters are required — `make(map)` in `maps1` and `map{}` in
`maps2` are incomplete types. The key type must be **comparable** (`==` must
work on it): strings, numbers, booleans, pointers, channels, interfaces, and
structs or arrays of comparable fields. Slices, maps, and functions cannot be
keys.

## 2. The zero value, and comma-ok

```go
phone := phoneBook["Anna"]        // "" — no such key, no error
phone, ok := phoneBook["Anna"]    // "", false
```

A read never fails. A missing key yields the value type's zero value, which is
indistinguishable from a stored zero unless you ask for the second result. That
is `maps3`: a typo'd key silently produces `""`, and the test failure is the
only signal.

```ascii
m := map[string]int{"a": 0}

m["a"]        -> 0        v, ok := m["a"]  -> 0, true   <- stored zero
m["missing"]  -> 0        v, ok := m["x"]  -> 0, false  <- absent
```

Use comma-ok whenever "absent" and "zero" mean different things — which is most
of the time for counters, flags, and caches. A `map[string]bool` used as a set
is the common exception: `if set[k]` reads correctly because absent and false
are the same answer.

## 3. Order is random on purpose

```go
for k, v := range m { … }   // different order every run
```

The runtime starts each iteration at a random bucket. This is not an
implementation accident you might get away with — it was added so that nobody
can depend on an order the implementation is free to change. If you need a
stable order, take the keys and sort them:

```go
keys := slices.Sorted(maps.Keys(m))   // Go 1.23
for _, k := range keys {
    use(k, m[k])
}
```

## 4. What a map is underneath

A map value is a **pointer to a runtime structure**, not the table itself. Two
consequences follow immediately:

- Passing a map to a function does **not** copy the contents. The function can
  insert and delete, and the caller sees it — unlike a slice, where appends do
  not propagate.
- The zero value of a map type is `nil`, and a `nil` map is **readable but not
  writable**: `m["k"]` returns the zero value, `m["k"] = 1` panics with
  `assignment to entry in nil map`. `make` or a literal before writing.

Go's implementation (a Swiss-table design since Go 1.24) stores entries in
groups and grows by rehashing when the load factor gets high — which is why
elements have no stable address: **you cannot take `&m[k]`**, and a map of
structs cannot have a field assigned in place (`m[k].n = 1` does not compile).
Store pointers, or read-modify-write the whole value.

Maps are also **not safe for concurrent use**. Concurrent writes are detected by
the runtime and abort the process with `fatal error: concurrent map writes` —
not a panic, not recoverable. Guard with a mutex, or see `safety2`.

## Gotchas

- **`make(map[K]V)` before writing.** A `nil` map panics on assignment.
- **A read of a missing key is not an error** — use comma-ok when it matters.
- **Iteration order is randomised**; sort the keys when output must be stable.
- **No `&m[k]`, and no assigning to a field of a struct value in a map.**
- **`len(m)` is the count; there is no `cap`.** `make(map[K]V, hint)` takes a
  size hint that only preallocates.
- **Concurrent writes crash the program**, hard. One writer, or a lock.
- **Deleting during `range` is safe** — the entry simply is not produced if not
  yet reached — but adding during `range` may or may not be visited.

## The exercises

- **maps1** — `make(map[string]int)`: both type parameters are part of the type.
- **maps2** — build the same map with a literal.
- **maps3** — a mistyped key returns `""`, not an error; fix the lookup and add
  the missing entry.

## Source references

- [Go blog: Go maps in action](https://go.dev/blog/maps) — comma-ok, sets, and
  ordering
- [Go spec: Map types](https://go.dev/ref/spec#Map_types) ·
  [Index expressions](https://go.dev/ref/spec#Index_expressions)
- [Go 1.24 release notes: Swiss-table maps](https://go.dev/doc/go1.24#performance)
- [pkg.go.dev: maps](https://pkg.go.dev/maps) — `Keys`, `Values`, `Clone`
- [pkg.go.dev: sync.Map](https://pkg.go.dev/sync#Map) — the concurrent case

**Next: [range](../range/) →** — the loop that walks all of these, and the copy
it hands you each iteration.
