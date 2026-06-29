// di3
// Constructor injection: a struct holds a dependency interface and delegates to it.

package main_test

import "testing"

type Store interface {
	Save(name string)
	Count() int
}

type Greeter struct {
	store Store
}

func (g Greeter) Greet(name string) string {
	g.store.Save(name)
	return "Hi, " + name
}

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
