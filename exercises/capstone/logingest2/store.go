// logingest2
// Stage two is storage. The service will accept events from many connections
// at once, so the thing holding them has to survive concurrent use: one
// goroutine appending while another reads must not corrupt the map, and a
// reader must not be handed a slice the store keeps mutating.

// I AM NOT DONE
package logingest

import (
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

// Add validates e and appends it to the events for e.Source. It returns the
// validation error unchanged when e is malformed, so callers can still match
// it with errors.Is(err, ErrInvalidEvent).
func (s *Store) Add(e Event) error {
	// FIXME: three things are wrong here.
	//
	//  1. e is never validated. Call e.Validate() and return its error.
	//  2. the map is written without holding s.mu. Take the write lock and
	//     release it with defer.
	//  3. nothing is actually appended.
	//
	// Writing a map from two goroutines at once is not a race you can get
	// away with — the runtime detects it and kills the program.
	return nil
}

// BySource returns the events recorded for source, oldest first. It returns an
// empty slice, not nil, when the source is unknown.
func (s *Store) BySource(source string) []Event {
	// FIXME: two things are wrong here.
	//
	//  1. the map is read without holding s.mu. Take the *read* lock — many
	//     readers may hold it at once, which is the point of RWMutex.
	//  2. returning s.events[source] hands the caller the store's own backing
	//     array. A later Add can overwrite what the caller is holding, and the
	//     caller can mutate the store from outside. Return a copy.
	//
	// slices.Clone is the copy you want. Remember to import "slices".
	return s.events[source]
}

// Len reports the total number of events across every source.
func (s *Store) Len() int {
	// FIXME: unlocked again, and it always reports zero. Take the read lock
	// and sum len() over the values.
	return 0
}
