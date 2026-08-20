# Sync primitives

"Share memory by communicating" is the advice, and it is good advice — but some
state genuinely is shared. A counter, a cache, a connection pool, a config
struct read by every request: passing those over a channel means routing every
access through one owner goroutine, which is often more machinery than the
problem deserves. For those cases the `sync` and `sync/atomic` packages let
several goroutines touch the same memory safely.

The three tools here cover the whole range. A `sync.Mutex` says *one goroutine at
a time*. A `sync.Once` says *this runs exactly once, everyone else waits for it*.
An atomic says *this single word updates indivisibly, no locking at all*. Pick by
the shape of the state, not by taste.

## 1. What a mutex actually is

`sync.Mutex` is two words: a state bitfield and a semaphore. The uncontended
path never enters the kernel — `Lock` is one compare-and-swap on the state word,
and `Unlock` is one atomic add. That is why a mutex around a few instructions is
cheap.

When the CAS fails, the loser spins briefly (a real optimisation on multicore:
the holder is often about to release), then parks on the semaphore. `Unlock`
wakes one waiter. If a waiter has been queued for more than **1 ms**, the mutex
flips into *starvation mode* and hands ownership directly to the head of the
queue instead of letting a freshly arriving goroutine barge in. That bounds the
tail latency — a goroutine cannot be starved indefinitely by a hot loop.

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```

Three habits that come out of the implementation:

- **The mutex goes next to the field it guards**, first in the struct, ideally
  with a comment naming what it covers. A mutex protects data, not code, and the
  code cannot tell you which data.
- **`defer Unlock` right after `Lock`.** Every return path, every panic,
  unlocks. The cost of `defer` is now negligible.
- **A `Mutex` is not reentrant.** Locking one you already hold deadlocks — the
  runtime has no idea who the owner is. A method that locks must not call
  another method that locks; factor out an unexported `…Locked` helper instead.

## 2. `Once`: exactly once, and everyone sees the result

```go
type Config struct {
    once sync.Once
    data map[string]string
}

func (c *Config) Load() {
    c.once.Do(c.load)   // 50 callers, one load
}
```

`Once` is a `done` flag plus a mutex. `Do` starts with an **atomic load** of the
flag — that is the entire cost after the first call, cheaper than taking a lock.
If the flag is unset, the slow path takes the mutex, re-checks the flag, runs
`f`, and sets the flag in a `defer`.

Two behaviours follow, and they are exactly what you want for lazy
initialisation:

- Callers that arrive during the first run **block until it finishes**. Nobody
  sees a half-built value, which a naive `if !inited { init() }` cannot promise.
- The flag is set in a `defer`, so **a panic inside `f` still marks it done**.
  `Once` means "attempted once", not "succeeded once". If the work can fail and
  should be retried, `Once` is the wrong tool — store the error and decide, or
  use a mutex you control.

Since Go 1.21 there are typed helpers when the point is producing a value:
`sync.OnceValue(f)` returns a function that computes `f` once and returns the
cached result; `sync.OnceFunc` and `sync.OnceValues` round it out.

## 3. Atomics: no lock at all

`sync/atomic` exposes single instructions the CPU already guarantees are
indivisible — load, store, add, swap, compare-and-swap on one word. There is no
queue and no parking, so an atomic add is roughly the cost of an ordinary
increment plus cache-line contention.

Prefer the **typed** wrappers introduced in Go 1.19:

```go
var hits atomic.Int64

hits.Add(1)
n := hits.Load()
```

They beat the older `atomic.AddInt64(&hits, 1)` functions on two counts: the
variable cannot accidentally be read non-atomically (`hits` has no usable raw
value), and the type is self-documenting. `atomic.Value` and the generic
`atomic.Pointer[T]` cover swapping whole structs by pointer.

The limit is the word. An atomic protects *one* variable; it cannot make two
fields consistent with each other. The moment your invariant spans two fields —
"`len` and `cap` must agree", "the map and the count must match" — you need a
mutex, because no CPU instruction updates both at once.

## 4. Choosing between them

| State | Tool | Why |
|---|---|---|
| One integer or flag | `atomic.Int64`, `atomic.Bool` | lock-free, no parking |
| One pointer swapped wholesale | `atomic.Pointer[T]` | readers never block |
| Two or more fields with an invariant | `sync.Mutex` | only a lock spans fields |
| Read-heavy, writes rare | `sync.RWMutex` | see `goroutine_safety` |
| One-time initialisation | `sync.Once` / `OnceValue` | blocks the herd, runs once |
| A value handed from producer to consumer | channel | ownership moves; nothing shared |

And the rule that makes all of it verifiable: **run concurrent tests with
`-race`**. `mise run test` does. The detector watches actual memory accesses, so
it finds the race you did not think about — and it is silent on correctly locked
code, which is the only proof you get that the lock is in the right place.

## Gotchas

- **Never copy a value containing a mutex.** `c := *counter` copies the lock in
  whatever state it was in. Pass pointers; `go vet` flags most copies.
- **A zero `Mutex`, `Once`, `RWMutex` and typed atomic is ready to use.** No
  constructor, no `New`. But once used, they must not move.
- **The race detector only sees code that runs.** A green `-race` run proves
  nothing about paths the test never took.
- **Contended atomics are not free.** Ten thousand goroutines hammering one
  `atomic.Int64` bounce a cache line between cores; sharded counters exist for
  a reason.
- **Locking around a channel send or an HTTP call** holds the lock for
  milliseconds and serialises everything. Compute under the lock, do I/O
  outside it.

## The exercises

- **sync1** — guard a struct's field with a `sync.Mutex`, `Lock`/`defer Unlock`,
  and watch `-race` go quiet.
- **sync2** — make lazy initialisation happen exactly once under 50 concurrent
  callers with `sync.Once`.
- **sync3** — replace `n++` with an atomic increment and see that a single word
  needs no lock at all.

## Source references

- [The Go Memory Model](https://go.dev/ref/mem) — the happens-before edges
  `Unlock`→`Lock` and `Do`→`Do` that make any of this valid
- [`src/sync/mutex.go`](https://github.com/golang/go/blob/master/src/sync/mutex.go)
  — the state bits, the spin, and the 1 ms starvation threshold
- [`src/sync/once.go`](https://github.com/golang/go/blob/master/src/sync/once.go)
  — the atomic fast path and the `defer`-set flag
- [pkg.go.dev: sync](https://pkg.go.dev/sync) ·
  [sync/atomic](https://pkg.go.dev/sync/atomic) ·
  [`OnceValue`](https://pkg.go.dev/sync#OnceValue)
- [Data Race Detector](https://go.dev/doc/articles/race_detector)

**Next: [context](../context/) →** — cancellation that crosses goroutine and API
boundaries, so the work itself stops instead of merely being ignored.
