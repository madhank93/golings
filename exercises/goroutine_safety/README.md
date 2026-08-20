# Goroutine safety

By this point in the series you can start goroutines, hand values between them,
wait on several at once, guard shared state, and cancel it all. What is left is
the discipline: the rules that decide *who* may touch a piece of state and *who*
is allowed to close a channel.

Those rules matter because the failure modes are so uneven. An unguarded integer
gives a quietly wrong number. An unguarded map gives a **fatal error that no
`recover` can catch**. Two goroutines closing the same channel gives a panic in
whichever loses the race — often only under load, in production. This chapter is
three concrete ownership decisions, each drilled under `go test -race`.

## 1. RWMutex: many readers, one writer

`sync.RWMutex` splits `Lock` into two modes. `RLock` may be held by any number of
readers at once; `Lock` is exclusive against everything. For read-heavy state —
a routing table read on every request and rebuilt once a minute — that turns a
serialisation point into a shared one.

```go
type Counter struct {
    mu sync.RWMutex
    n  int
}

func (c *Counter) Inc()       { c.mu.Lock();  defer c.mu.Unlock();  c.n++ }
func (c *Counter) Value() int { c.mu.RLock(); defer c.mu.RUnlock(); return c.n }
```

Inside, the lock keeps a reader count and a writer semaphore. A waiting writer
immediately blocks *new* readers (by pushing the reader count negative) and then
waits for the readers already inside to leave — so a steady stream of readers
cannot starve a writer.

That write-preferring design has one sharp edge: **read locks are not
recursive**. If a reader takes `RLock`, a writer queues behind it, and that same
reader then takes `RLock` again, the second acquisition waits behind the writer
while the writer waits behind the first — a three-way deadlock in code that looks
like two harmless reads.

Do not reach for `RWMutex` by default. It has more state to maintain than a
`Mutex`, so under short critical sections and a balanced read/write mix a plain
`Mutex` is faster. Switch when reads dominate *and* the critical section is long
enough to matter — measure with `go test -bench` before assuming.

## 2. sync.Map, and why it is a niche tool

A plain map is not safe for concurrent use, and the runtime does not merely
corrupt it — it detects the situation and kills the process:

```
fatal error: concurrent map writes
```

Not a panic. Not recoverable. The map implementation writes across several
words (bucket, tombstones, count, growth state) and cannot be left half-updated,
so the runtime aborts.

The default fix is a `map` plus a `Mutex` or `RWMutex` — obvious, typed, and fast.
`sync.Map` is the specialist:

```go
var m sync.Map
m.Store("api", 1)
v, ok := m.Load("api")     // (any, bool) — assert the type back
id := v.(int)
```

Reach for it in the two cases its own documentation names: a key written once
and read many times, or disjoint sets of keys touched by disjoint goroutines.
Its implementation is built for exactly those — historically a lock-free read
map with a mutex-guarded dirty map behind it, and since Go 1.24 a concurrent
hash-trie that scales better under mixed load.

The costs are real, though: keys and values are `any`, so you lose type safety
and pay an allocation to box them, `Range` sees a snapshot rather than a
consistent view, and there is no `len`. A generic `map[K]V` behind an `RWMutex`
is usually the better default; `safety2` exists so you know what `sync.Map`
buys when you do choose it.

## 3. Channel ownership: the owner closes, once

Closing is a statement about **sends**: "no more values will ever be sent". Only
the goroutine that sends can honestly make that claim, so:

> The goroutine that owns a channel closes it — exactly once. Receivers never
> close.

Break the rule and you get one of two panics — `close of closed channel` when two
owners race, or `send on closed channel` when a receiver closes out from under a
producer. `safety3` is the first case, made deterministic: three workers range
over one channel, and each closing it means the second one dies.

The structural fix is directional types, from the `channels` chapter. Hand
consumers a `<-chan T` and closing it does not compile:

```go
func consume(in <-chan int) {   // cannot close(in) — receive-only
    for v := range in { … }
}
```

When several producers share one output, nobody owns it individually — so hoist
ownership to a closer goroutine that waits for all of them, exactly as fan-in
does:

```go
go func() { wg.Wait(); close(out) }()
```

And when the *consumer* needs to stop the producers, do not close the data
channel from the receiving end. Close a separate `done` channel, or cancel a
`context`, and let the owner do the closing.

## 4. Leak-free by construction

Every goroutine you start must have an answer to "what makes it return?". Three
that work:

- its input channel gets closed (pipeline stages),
- it selects on `ctx.Done()` alongside its work (anything that blocks on I/O),
- it finishes a bounded loop (a worker over a fixed batch).

Anything else is a leak: a parked goroutine holds its stack and everything it
references, and it never appears in a test failure. Two tools make leaks
visible — `runtime.NumGoroutine()` before and after in a test, and
`go.uber.org/goleak` as a `TestMain` check. The `profiling` chapter's goroutine
profile shows the stacks in a live process.

## Gotchas

- **`fatal error: concurrent map writes` is not catchable.** No `recover`, no
  graceful degradation. Guard the map or use `sync.Map`.
- **The race detector finds races, not deadlocks or leaks.** Different tools:
  `-race`, then goroutine counts, then the goroutine profile.
- **`RLock` is not reentrant, and neither is `Lock`.**
- **Copying a struct that holds a mutex or a `sync.Map`** silently gives you two
  independent locks. `go vet` catches the obvious ones.
- **`sync.Map` zero value is ready**, but `Load` returns `any` — a wrong type
  assertion panics at the call site, not at `Store`.
- **A `for range ch` in a worker is not "drain what is there"**; it blocks until
  the channel is closed. A worker pool over a channel nobody closes never
  finishes.

## The exercises

- **safety1** — add a `sync.RWMutex` to a struct read and written by 100
  goroutines, and pick `Lock` vs `RLock` per method.
- **safety2** — replace a plain map that crashes under concurrent writes with a
  `sync.Map`, and type-assert the values back.
- **safety3** — remove the `close` from a worker that only receives, leaving the
  owner as the single closer.

## Source references

- [The Go Memory Model](https://go.dev/ref/mem)
- [Data Race Detector](https://go.dev/doc/articles/race_detector) — how `-race`
  instruments memory and reads its reports
- [`src/sync/rwmutex.go`](https://github.com/golang/go/blob/master/src/sync/rwmutex.go)
  — reader count, writer preference, and the recursive-RLock deadlock note
- [pkg.go.dev: sync.Map](https://pkg.go.dev/sync#Map) — the two cases it is built
  for, in its own words
- [Go 1.24 release notes: sync.Map](https://go.dev/doc/go1.24#sync) — the
  hash-trie rewrite
- [go.uber.org/goleak](https://pkg.go.dev/go.uber.org/goleak) — failing a test
  on leaked goroutines

**Next: [synctest](../synctest/) →** — testing all of this deterministically,
without `time.Sleep` and without flakes.
