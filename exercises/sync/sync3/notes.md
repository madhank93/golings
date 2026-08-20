## sync3 — lock-free counters with sync/atomic

```go
var hits int64
atomic.AddInt64(&hits, 1)   // atomic increment
atomic.LoadInt64(&hits)     // atomic read
```

**Why it works**

- `atomic.AddInt64` performs the read-add-write as **one indivisible CPU
  operation**, so 1000 concurrent increments never lose an update — no mutex
  needed.

**Under the hood**

- This compiles to a single instruction the CPU guarantees is atomic (`LOCK
  XADD` on x86, `LDADD` on arm64). There is no queue and no parking, so an
  uncontended atomic costs about as much as a plain increment. Under heavy
  contention the cost is cache-line ping-pong between cores rather than
  scheduling — which is why very hot counters are sometimes sharded.

**Common mistake**

- Mixing atomic and plain access to the same variable: `atomic.AddInt64(&hits,1)`
  in one goroutine and `hits++` or a bare `hits` read in another is still a data
  race. *Every* access has to go through the atomic API. Prefer the Go 1.19 typed
  wrappers (`var hits atomic.Int64`; `hits.Add(1)`, `hits.Load()`) — they make
  the non-atomic access impossible to write.

**Key detail:** atomics only cover **single-word** operations — one counter, one
flag, one pointer. The moment two fields must stay consistent with each other, no
instruction can update both, and you need a mutex. `atomic.Pointer[T]` covers
swapping a whole struct by pointer, which is how read-mostly config is usually
published.

**See also:** sync1 (mutex when the invariant spans fields) · concurrent2 (the
race being fixed) · safety3 (`atomic.Int64` in a worker) · the
[sync chapter](../README.md)

**References**

- pkg.go.dev — sync/atomic: https://pkg.go.dev/sync/atomic ·
  atomic.Int64: https://pkg.go.dev/sync/atomic#Int64
- The Go Memory Model — atomics: https://go.dev/ref/mem#atomic
- Go by Example — Atomic Counters: https://gobyexample.com/atomic-counters
