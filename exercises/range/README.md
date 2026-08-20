# Range

Go has exactly one loop keyword, `for`, and `range` is the clause that walks a
collection with it. One form covers slices, arrays, strings, maps, channels,
integers (Go 1.22), and functions (Go 1.23) — with the pair of values it yields
depending on what you hand it.

The subtlety worth carrying away is that `range` gives you a **copy** of each
element, and evaluates its subject **once**, up front. Both are why the loop
sometimes does not do what a first reading suggests.

## 1. What each subject yields

| Range over | First value | Second value |
|---|---|---|
| slice, array | index | copy of the element |
| string | **byte** offset | `rune` (decoded UTF-8) |
| map | key | value |
| channel | received value | — (one value only) |
| `int` (1.22) | 0 … n-1 | — |
| `func` iterator (1.23) | whatever it yields | optional second |

```go
for i, v := range nums { … }   // both
for i := range nums    { … }   // index only
for _, v := range nums { … }   // value only — range1
for range nums         { … }   // neither: just repeat len(nums) times
for i := range 5       { … }   // 0,1,2,3,4 — no slice needed (Go 1.22)
```

`range2` is the map form, where the two values are key and value rather than
index and element. `range3` is the filter shape — range the input, test each
value, `append` the ones that pass to a new slice.

## 2. The value is a copy

```go
type Item struct{ n int }
items := []Item{{1}, {2}}

for _, it := range items {
    it.n = 99          // writes to the copy; items is unchanged
}
for i := range items {
    items[i].n = 99    // writes through the index — this works
}
```

This is the single most common `range` surprise. The loop variable is assigned a
copy of each element, so mutating it changes nothing. Index when you need to
write; range the value when you only need to read. For large structs the copy
also costs — ranging the index avoids it.

```ascii
for _, v := range items      for i := range items

  items[0] ──copy──► v         items[0] ◄── items[i].n = 99
  (writing v changes nothing)  (writes the real element)
```

## 3. The range expression is evaluated once

```go
for i, v := range s {
    s = append(s, v)   // does NOT loop forever
}
```

The length is captured before the first iteration, so appending inside the loop
cannot extend it. The same is true of the slice header: `range s` walks the
slice as it was when the loop started, even if `s` is reassigned in the body.

For maps, only the *start* is fixed — entries added during iteration may or may
not be visited, and entries deleted before they are reached will not be.

## 4. The loop variable, then and now

Before **Go 1.22**, one loop variable was shared by every iteration. The classic
bug:

```go
for _, v := range items {
    go func() { use(v) }()   // pre-1.22: every goroutine saw the LAST v
}
```

Since 1.22 each iteration gets a **fresh** variable, so closures and goroutines
capture what you expect and the `v := v` shadow-copy line is no longer needed.
This repo is on Go 1.26 — write the plain closure. When reading older code (or
older blog posts), remember the semantics changed under it.

## 5. Two newer forms

```go
for i := range 10 { … }                  // Go 1.22: range over an int

for k, v := range myIterator { … }        // Go 1.23: range over func
```

The second is the iterator protocol: any `func(yield func(K, V) bool)` can be
ranged directly, which is what `iter.Seq` and `maps.Keys` produce. The
`iterators` chapter covers it.

Ranging a **channel** is the other important form: it receives until the channel
is closed, which is how consumers in the concurrency chapters drain a producer.

## Gotchas

- **The value is a copy** — mutate via `items[i]`, not the loop variable.
- **`range` over a string yields byte offsets**, so the index jumps for
  multi-byte characters. `len(s)` is bytes, not iterations.
- **Map order is randomised** on every run. Sort the keys for stable output.
- **`for range ch` needs someone to `close(ch)`**, or it blocks forever.
- **Appending to the slice you are ranging** does not extend the loop.
- **`for i := range 0`** runs zero times, as does ranging a `nil` slice or map —
  no guard needed.
- **`break` and `continue`** work as expected; a labelled `break` exits an outer
  loop.

## The exercises

- **range1** — range a slice for its values, ignoring the index with `_`.
- **range2** — range a map for keys and values.
- **range3** — use `range` to filter: collect the even numbers into a new slice.

## Source references

- [Go spec: For statements with range clause](https://go.dev/ref/spec#For_range)
- [Go 1.22 release notes: loop variable scoping](https://go.dev/doc/go1.22#language) ·
  [Go blog: Fixing for loops in Go 1.22](https://go.dev/blog/loopvar-preview)
- [Go 1.23 release notes: range over function](https://go.dev/doc/go1.23#language)
- [Go by Example: Range over Built-in Types](https://gobyexample.com/range-over-built-in-types)

**End of the Collections tier.** Next: [structs](../structs/) — giving these
values names and behaviour instead of positions.
