# Modern Go (1.21 – 1.25)

Go changes slowly, but the releases since 1.21 changed enough of the everyday
vocabulary that code written from memory five years ago now reads as dated. This
chapter is the shortlist: the built-ins you no longer write helpers for, the two
loop changes, and the concurrency helper that removes a boilerplate people got
wrong.

Everything here is language-level or standard-library — nothing to import,
nothing to install.

## 1. `min`, `max`, `clear` (Go 1.21)

```go
lo := min(3, 1, 2)     // 1  — two or more args, any ordered type
hi := max(3, 1, 2)     // 3
clear(m)               // empties a map in place
clear(s)               // zeroes every element of a slice (keeps its length)
```

`min` and `max` are **built-ins**, not functions from a package: no import, and
they work for any ordered type including strings. They replace the
hand-written `func maxInt(a, b int) int` that used to appear in every codebase.

`clear` is the one with two behaviours, and the difference matters: on a **map**
it deletes every entry (the map stays usable, `len` becomes 0); on a **slice**
it sets every element to the zero value and the **length is unchanged**. To
empty a slice you still write `s = s[:0]`.

## 2. Range over an integer (Go 1.22)

```go
for i := range 5 { … }     // i = 0,1,2,3,4
for range 3 { … }          // just repeat three times
```

`modern2`. The three-clause `for i := 0; i < n; i++` still exists and is still
right when you need a step or a different start, but the plain count-up loop —
by far the most common — no longer needs it. Ranging a negative or zero value
runs zero times.

## 3. Per-iteration loop variables (Go 1.22)

This is the change with the largest blast radius. Before 1.22, a loop declared
**one** variable reused by every iteration:

```ascii
pre-1.22                        1.22 and later

  item ──┬── closure 1            item#1 ── closure 1
         ├── closure 2            item#2 ── closure 2
         └── closure 3            item#3 ── closure 3
  one variable, everyone          each iteration declares its own
  sees the last value
```

So closures and goroutines created in a loop all captured the same variable and
usually saw the final value — the single most reported Go bug, worked around
with `item := item` at the top of the body. Since 1.22 each iteration declares a
fresh variable and the plain closure is correct.

The subtlety `modern5` drills: this only applies to variables the **loop
declares**. Hoist the variable out and assign it with `=` and you are back to
the old, shared-variable behaviour — one variable, every closure seeing the last
value. The fix is to let the loop declare it with `:=`.

## 4. Range over a function (Go 1.23)

```go
func countUp(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := 1; i <= n; i++ {
            if !yield(i) {     // consumer broke out — stop producing
                return
            }
        }
    }
}

for v := range countUp(3) { … }     // 1, 2, 3
```

`modern3`. Any function shaped `func(yield func(V) bool)` can be ranged
directly. `yield` returns `false` when the consumer stopped (a `break`, a
`return`, an error path), and an iterator that ignores that return value keeps
producing into a loop that has already left. The `iterators` chapter covers the
protocol in full.

## 5. `WaitGroup.Go` (Go 1.25)

```go
var wg sync.WaitGroup
for _, n := range nums {
    wg.Go(func() {          // Add(1), go, and defer Done() — in one call
        total.Add(int64(n))
    })
}
wg.Wait()
```

`modern4`. The classic three-line dance — `wg.Add(1)`, `go func(){`,
`defer wg.Done()` — had two failure modes that this removes: forgetting `Add`
entirely (so `Wait` returns immediately and you read results that do not exist
yet, which is the bug in the exercise), and calling `Add` *inside* the goroutine,
which races with `Wait`.

## Also worth knowing

- **`for i, v := range slices.All(s)`** and friends (1.23) — iterator versions
  of the collection helpers.
- **`errors.Join`** (1.20) — several failures as one error (`errors5`).
- **`sync.OnceValue` / `OnceFunc`** (1.21) — lazy init that returns a value.
- **`slices` and `maps`** (1.21) — `Sort`, `Contains`, `Clone`, `Keys`, `Values`;
  the generic helpers everyone used to hand-roll.
- **`testing/synctest`** (GA 1.25) — deterministic time in tests (`synctest1`).
- **`omitzero` in JSON tags** (1.24) — the `omitempty` that actually means
  "zero" (`stdlib7`).

## Gotchas

- **`clear(slice)` does not shorten it.** Use `s = s[:0]` to empty a slice.
- **`min`/`max` need at least two arguments** and all of one ordered type.
- **The 1.22 loop fix applies only to loop-declared variables** — `modern5`'s
  whole point.
- **An iterator that ignores `yield`'s `false`** keeps running after the consumer
  has left.
- **`wg.Go` still needs `wg.Wait()`**; it replaces `Add`/`Done`, not the wait.
- **Language changes are gated by the `go` line in `go.mod`.** A module
  declaring `go 1.21` does not get 1.22 loop semantics even on a new toolchain.

## The exercises

- **modern1** — replace hand-rolled helpers with built-in `min`, `max`, `clear`.
- **modern2** — range over an integer instead of a three-clause loop.
- **modern3** — implement an `iter.Seq[int]` body, respecting `yield`'s result.
- **modern4** — launch tracked goroutines with `wg.Go`.
- **modern5** — let the loop declare the variable so each closure captures its
  own.

## Source references

- [Go 1.21 release notes](https://go.dev/doc/go1.21#language) ·
  [1.22](https://go.dev/doc/go1.22#language) ·
  [1.23](https://go.dev/doc/go1.23#language) ·
  [1.24](https://go.dev/doc/go1.24) · [1.25](https://go.dev/doc/go1.25)
- [Go blog: Fixing for loops in Go 1.22](https://go.dev/blog/loopvar-preview)
- [Go blog: Range Over Function Types](https://go.dev/blog/range-functions)
- [pkg.go.dev: builtin](https://pkg.go.dev/builtin) ·
  [sync.WaitGroup.Go](https://pkg.go.dev/sync#WaitGroup.Go)

**Next: [iterators](../iterators/) →** — the `iter.Seq` protocol from
`modern3`, in full, including how to pause one.
