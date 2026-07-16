// safety2
// sync.Map is a concurrent map built for stable keys and many reads. Its
// Store/Load methods are safe for simultaneous use — a plain map is not, and
// concurrent writes to one crash with "fatal error: concurrent map writes".

package main_test

import (
	"strconv"
	"sync"
	"testing"
)

// Registry maps handler names to ids, populated concurrently at startup.
type Registry struct {
	m sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Set(name string, id int) {
	r.m.Store(name, id)
}

func (r *Registry) Get(name string) (int, bool) {
	v, ok := r.m.Load(name)
	if !ok {
		return 0, false
	}
	return v.(int), true
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
