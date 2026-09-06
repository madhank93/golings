// logingest4
// Stage four puts a worker pool between the handler and the store. Ingest
// should not block on storage, and a burst of connections should not turn into
// a burst of goroutines — a fixed set of workers draining one channel gives you
// backpressure instead.

// I AM NOT DONE
package logingest

import (
	"sync"
	"sync/atomic"
)

// Pool fans events out to a fixed number of workers, each of which writes to
// the same Store. Submit is safe to call from many goroutines at once.
//
// The zero value is not usable; call NewPool.
type Pool struct {
	store    *Store
	events   chan Event
	wg       sync.WaitGroup
	rejected atomic.Int64
}

// NewPool returns a Pool with workers goroutines already running, each draining
// the shared channel until it is closed.
func NewPool(store *Store, workers int) *Pool {
	p := &Pool{
		store: store,
		// A small buffer lets a producer get ahead of the workers without
		// blocking on every single send.
		events: make(chan Event, workers),
	}
	// FIXME: start `workers` goroutines. Each one ranges over p.events and
	// calls p.store.Add for every event, counting a rejected event with
	// p.rejected.Add(1). Track them with p.wg so Close can wait:
	//
	//	for range workers {
	//		p.wg.Go(func() { ... })
	//	}
	//
	// sync.WaitGroup.Go (Go 1.25) does the Add(1)/defer Done() pairing for
	// you. Right now no worker exists, so nothing is ever stored.
	return p
}

// Submit hands e to the pool. It blocks while every worker is busy and the
// buffer is full, which is the backpressure that keeps a burst of requests from
// becoming unbounded memory.
func (p *Pool) Submit(e Event) {
	// FIXME: send e on p.events.
	// Right now the event is dropped on the floor.
}

// Close stops the pool and waits for the workers to finish the events already
// submitted. It must be called exactly once, after the last Submit.
func (p *Pool) Close() {
	// FIXME: close p.events, then p.wg.Wait().
	//
	// Order matters and so does the pairing. Closing is what ends each
	// worker's `range`; without it Wait blocks forever. Waiting is what makes
	// Close a real barrier; without it Close returns while events are still
	// in flight and the caller reads a half-filled store.
}

// Rejected reports how many submitted events failed validation.
func (p *Pool) Rejected() int64 {
	// FIXME: return p.rejected.Load().
	//
	// atomic.Int64 rather than a plain int64 with a mutex: every worker
	// increments this, and a counter is the one case where the atomic is
	// simpler than the lock.
	return 0
}
