// di3
// Constructor injection: a struct holds a dependency interface and delegates to it.

// I AM NOT DONE
package main_test

import "testing"

// Store abstracts where greeted names are recorded.
type Store interface {
	Save(name string)
	Count() int
}

// Greeter depends on a Store, injected when the Greeter is built.
type Greeter struct {
	store Store
}

// Greet records the name in the store and returns "Hi, <name>".
func (g Greeter) Greet(name string) string {
	// FIXME: record the name via the injected store, then return the greeting.
	return ""
}

// memStore is an in-memory test double for Store.
type memStore struct{ names []string }

func (m *memStore) Save(name string) { m.names = append(m.names, name) }
func (m *memStore) Count() int       { return len(m.names) }

func TestGreeterUsesStore(t *testing.T) {
	store := &memStore{}
	g := Greeter{store: store}
	if got := g.Greet("Go"); got != "Hi, Go" {
		t.Errorf("want %q, got %q", "Hi, Go", got)
	}
	if store.Count() != 1 {
		t.Errorf("want 1 saved name, got %d", store.Count())
	}
}
