package logingest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func poolEvent(source, msg string) Event {
	return Event{
		Source:  source,
		Level:   "info",
		Message: msg,
		At:      time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC),
	}
}

// Close is a barrier: everything submitted before it must be in the store after
// it. A Close that forgets to Wait fails here intermittently, which is exactly
// why the count is checked immediately after Close returns.
func TestPoolStoresEverythingSubmittedBeforeClose(t *testing.T) {
	store := NewStore()
	p := NewPool(store, 4)

	const n = 200
	for i := range n {
		p.Submit(poolEvent("api", fmt.Sprintf("msg %d", i)))
	}
	p.Close()

	if got := store.Len(); got != n {
		t.Errorf("store.Len() = %d, want %d", got, n)
	}
	if got := p.Rejected(); got != 0 {
		t.Errorf("Rejected() = %d, want 0", got)
	}
}

// Invalid events are counted, not stored, and must not kill a worker.
func TestPoolCountsRejectedEvents(t *testing.T) {
	store := NewStore()
	p := NewPool(store, 2)

	p.Submit(poolEvent("api", "good"))
	p.Submit(poolEvent("", "no source"))
	p.Submit(poolEvent("api", "still good"))
	p.Close()

	if got := store.Len(); got != 2 {
		t.Errorf("store.Len() = %d, want 2", got)
	}
	if got := p.Rejected(); got != 1 {
		t.Errorf("Rejected() = %d, want 1", got)
	}
}

// Submit is called from many goroutines at once, which is how the HTTP handler
// will use it. Run with -race.
func TestPoolAcceptsConcurrentSubmits(t *testing.T) {
	store := NewStore()
	p := NewPool(store, 4)

	const producers, each = 8, 25
	var wg sync.WaitGroup
	for prod := range producers {
		wg.Go(func() {
			for i := range each {
				p.Submit(poolEvent(fmt.Sprintf("src-%d", prod), fmt.Sprintf("msg %d", i)))
			}
		})
	}
	wg.Wait()
	p.Close()

	if got := store.Len(); got != producers*each {
		t.Errorf("store.Len() = %d, want %d", got, producers*each)
	}
}

// A pool with more workers than events must still terminate. This is the test
// that catches a Close which never closes the channel: the workers block on an
// empty range forever and Wait deadlocks.
func TestPoolCloseTerminatesWithIdleWorkers(t *testing.T) {
	store := NewStore()
	p := NewPool(store, 8)
	p.Submit(poolEvent("api", "only one"))

	done := make(chan struct{})
	go func() {
		p.Close()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Close did not return: the events channel is never closed, so the workers never finish")
	}

	if got := store.Len(); got != 1 {
		t.Errorf("store.Len() = %d, want 1", got)
	}
}
