## pprof2 — measuring retention with MemStats

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

- `HeapAlloc` is the number of live heap bytes. Reading it before and after,
  with a forced GC on each side, gives the bytes the allocation **retains** —
  not the bytes it churned through.

**Under the hood**

- Both extra calls are load-bearing. Without `runtime.GC()` the reading includes
  garbage not yet collected, so you measure allocation rate instead of
  retention. Without `runtime.KeepAlive(keep)` the compiler may treat the value
  as dead before the second reading and report zero — the value is still in scope
  in the source, but scope is not what keeps things alive.

**Common mistake**

- Calling `ReadMemStats` in a request path. It **stops the world** to take a
  consistent snapshot; it is a diagnostic, not a metric to sample continuously.

**Key detail:** the useful fields are `HeapAlloc` (live now), `TotalAlloc`
(cumulative, never decreases), `Sys` (from the OS), `NumGC` and `PauseTotalNs`.
For allocation **sites** rather than totals, use the heap profile:
`go test -memprofile mem.out`, then `-alloc_space` vs `-inuse_space`.

**See also:** pprof1 (CPU) · pprof3 · unsafe1 (struct size) · testadv4 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — runtime.MemStats: https://pkg.go.dev/runtime#MemStats ·
  runtime.KeepAlive: https://pkg.go.dev/runtime#KeepAlive
- Go — Diagnostics: https://go.dev/doc/diagnostics
