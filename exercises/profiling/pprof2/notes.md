## pprof2 — measure retained heap with MemStats

```go
runtime.GC()
runtime.ReadMemStats(&before)
keep := alloc()
runtime.GC()
runtime.ReadMemStats(&after)
runtime.KeepAlive(keep)
return int64(after.HeapAlloc) - int64(before.HeapAlloc)
```

**Why it works**

- `runtime.ReadMemStats` snapshots the heap. Reading `HeapAlloc` **before** and
  **after** — with a `runtime.GC()` each time so freed garbage isn't counted —
  gives the bytes the allocation **retains**.

**Key detail:** two subtleties. `runtime.GC()` before the "after" read forces
collection so only still-live memory is measured. And `runtime.KeepAlive(keep)`
stops the optimizer from freeing `keep` early — without it the compiler might
decide the allocation is dead before you measure it.

**uint64 wrap hazard:** `HeapAlloc` is a `uint64`, so `after - before` on a heap
that *shrank* wraps to a number near 2⁶⁴ instead of going negative. A naive
`if got < want` check then passes on a measurement that is pure nonsense. Convert
both sides to `int64` first so a shrinking heap reads as a negative delta.

**Why a tolerance band, not an exact count:** the result is a *net* heap delta.
The second `runtime.GC()` also frees garbage that was still live at the first
snapshot, so measured growth lands a few KB under the true allocation size — a
1 MiB slice typically reports ~1008–1043 KB. Exact-byte assertions on MemStats
are flaky by construction; assert a range.

**References**

- pkg.go.dev — runtime.ReadMemStats: https://pkg.go.dev/runtime#ReadMemStats
