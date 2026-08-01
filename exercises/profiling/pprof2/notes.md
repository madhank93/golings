## pprof2 — measure retained heap with MemStats

```go
runtime.GC()
runtime.ReadMemStats(&before)
keep := alloc()
runtime.GC()
runtime.ReadMemStats(&after)
runtime.KeepAlive(keep)
return after.HeapAlloc - before.HeapAlloc
```

**Why it works**

- `runtime.ReadMemStats` snapshots the heap. Reading `HeapAlloc` **before** and
  **after** — with a `runtime.GC()` each time so freed garbage isn't counted —
  gives the bytes the allocation **retains**.

**Key detail:** two subtleties. `runtime.GC()` before the "after" read forces
collection so only still-live memory is measured. And `runtime.KeepAlive(keep)`
stops the optimizer from freeing `keep` early — without it the compiler might
decide the allocation is dead before you measure it.

**References**

- pkg.go.dev — runtime.ReadMemStats: https://pkg.go.dev/runtime#ReadMemStats
