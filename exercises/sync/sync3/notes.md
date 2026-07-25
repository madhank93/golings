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

**Nuance:** atomics are faster than a mutex but only cover **single-word**
operations (one counter, one flag, one pointer). The moment you must update two
fields together, you need a mutex. Always go through `atomic.Load`/`Store` for
*every* access to the value — mixing a plain read with atomic writes is still a
race. (Go 1.19+ also offers `atomic.Int64` with methods.)

**References**

- Go by Example — Atomic Counters: https://gobyexample.com/atomic-counters
- sync/atomic: https://pkg.go.dev/sync/atomic
