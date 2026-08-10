// Carried forward, already solved. Read it, do not edit it.
package logingest

import (
	"slices"
	"sync"
)

// Store keeps events grouped by source. It is safe for concurrent use.
//
// The zero value is not usable; call NewStore.
type Store struct {
	mu     sync.RWMutex
	events map[string][]Event
}

// NewStore returns an empty, ready-to-use Store.
func NewStore() *Store {
	return &Store{events: make(map[string][]Event)}
}

// Add validates e and appends it to the events for e.Source.
func (s *Store) Add(e Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[e.Source] = append(s.events[e.Source], e)
	return nil
}

// BySource returns a copy of the events recorded for source, oldest first.
func (s *Store) BySource(source string) []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if got := s.events[source]; got != nil {
		return slices.Clone(got)
	}
	return []Event{}
}

// Len reports the total number of events across every source.
func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	for _, evs := range s.events {
		n += len(evs)
	}
	return n
}
