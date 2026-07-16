// safety2
// sync.Map is a concurrent map built for stable keys and many reads. Its
// Store/Load methods are safe for simultaneous use — a plain map is not, and
// concurrent writes to one crash with "fatal error: concurrent map writes".

// I AM NOT DONE
package main_test

import (
	"strconv"
	"sync"
	"testing"
)

// Registry maps handler names to ids, populated concurrently at startup.
type Registry struct {
	// FIXME: a plain map is not safe for concurrent writes. Replace m with a
	// sync.Map and rewrite Set/Get to use m.Store and m.Load (Load returns
	// (any, bool), so type-assert the value to int).
	m map[string]int
}

func NewRegistry() *Registry {
	return &Registry{m: map[string]int{}}
}

func (r *Registry) Set(name string, id int) {
	r.m[name] = id
}

func (r *Registry) Get(name string) (int, bool) {
	v, ok := r.m[name]
	return v, ok
}

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	var wg sync.WaitGroup
	for i := range 50 {
		wg.Go(func() {
			r.Set(strconv.Itoa(i), i)
		})
	}
	wg.Wait()
	if v, ok := r.Get("7"); !ok || v != 7 {
		t.Errorf("want 7,true; got %d,%v", v, ok)
	}
}
