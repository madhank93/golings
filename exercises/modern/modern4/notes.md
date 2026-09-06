## modern4 — WaitGroup.Go

```go
for _, n := range nums {
    wg.Go(func() {
        total.Add(int64(n))
    })
}
wg.Wait()
```

**Why it works**

- `wg.Go(f)` (Go 1.25) does the `Add(1)`, starts the goroutine, and defers
  `Done()` — the three-line dance in one call. The broken version uses a bare
  `go` with no `Add`, so the counter never rises, `Wait` returns immediately, and
  `total` is read before the goroutines have run.

**Under the hood**

- The two classic failure modes it removes are exactly those: **no `Add` at all**
  (Wait returns early, as here), and **`Add` inside the goroutine**, which races
  with `Wait` — the counter may still be 0 when `Wait` checks it.

**Common mistake**

- Thinking `wg.Go` replaces the wait too. It does not: `wg.Wait()` is still what
  blocks until the count returns to zero.

**Key detail:** the accumulator is an `atomic.Int64` for a reason — several
goroutines add to it concurrently, and a plain `int64` would be a data race that
`go test -race` reports. `wg.Go` handles the counting, not the sharing.

**See also:** concurrent1 (the long form) · sync3 (atomics) · safety1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — sync.WaitGroup.Go: https://pkg.go.dev/sync#WaitGroup.Go
- Go 1.25 release notes: https://go.dev/doc/go1.25
