// logingest4
// Stage four puts a worker pool between the handler and the store. Ingest
// should not block on storage, and a burst of connections should not turn into
// a burst of goroutines — a fixed set of workers draining one channel gives you
// backpressure instead.

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
		store:  store,
		events: make(chan Event, workers),
	}
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

// Submit hands e to the pool, blocking while every worker is busy and the
// buffer is full.
func (p *Pool) Submit(e Event) {
	p.events <- e
}

// Close stops the pool and waits for the workers to drain what was submitted.
func (p *Pool) Close() {
	close(p.events)
	p.wg.Wait()
}

// Rejected reports how many submitted events failed validation.
func (p *Pool) Rejected() int64 {
	return p.rejected.Load()
}
