## logingest4 — a worker pool, and the two halves of Close

```go
func NewPool(store *Store, workers int) *Pool {
	p := &Pool{store: store, events: make(chan Event, workers)}
	for range workers {
		p.wg.Go(func() {
			for e := range p.events {
				if err := p.store.Add(e); err != nil {
					p.rejected.Add(1)
				}
			}
		})
	}
	return p
}

func (p *Pool) Submit(e Event) { p.events <- e }

func (p *Pool) Close() {
	close(p.events)
	p.wg.Wait()
}

func (p *Pool) Rejected() int64 { return p.rejected.Load() }
```

**Why it works**

- `for e := range ch` on a channel ends **by itself** when the channel is
  closed and drained. That is the whole shutdown protocol — no quit channel, no
  sentinel value, no counter of expected events.
- A fixed `workers` count is the point of a pool. Spawning a goroutine per
  request works right up until a traffic spike, at which point you have a
  million goroutines and no memory. Here a spike blocks the producer instead.
- The buffer gives a producer a little room to run ahead, but it is bounded, so
  a sustained overload still pushes back on the caller rather than queueing
  without limit. Backpressure is a feature.
- `sync.WaitGroup.Go` (Go 1.25) replaces the `wg.Add(1)` / `defer wg.Done()`
  pair. The old form is still correct; the new one cannot be got wrong.

**Key detail:** `Close` has two halves and needs both. `close(p.events)` is what
ends each worker's `range` — leave it out and `Wait` blocks forever, which the
idle-workers test catches with an explicit timeout rather than by hanging the
suite. `p.wg.Wait()` is what makes `Close` a barrier — leave *it* out and `Close`
returns while workers are still writing, so the caller reads a half-filled store
and gets a different answer on every run. That second bug is the nastier one: it
usually passes locally and fails in CI.

**Key detail:** closing a channel is the *sender's* job, and only once. A worker
must never close `p.events`, and `Close` must not be called twice — a second
`close` panics with "close of closed channel". This is why `Close` is documented
as "call exactly once, after the last Submit" rather than being made idempotent:
the discipline belongs at the call site, where the program knows when the last
event has been submitted.

**Key detail:** `atomic.Int64` for the rejected counter, not a mutex. Every
worker increments it and one reader reads it — the single case where the atomic
is genuinely simpler than the lock. Note it is a field of the struct, not a
pointer: `atomic.Int64` has its own methods and must not be copied, which is
another reason `NewPool` returns `*Pool`.

**References**

- Go blog — Go Concurrency Patterns: Pipelines: https://go.dev/blog/pipelines
- pkg.go.dev — sync.WaitGroup.Go: https://pkg.go.dev/sync#WaitGroup.Go
- pkg.go.dev — sync/atomic.Int64: https://pkg.go.dev/sync/atomic#Int64
- Go spec — close: https://go.dev/ref/spec#Close
