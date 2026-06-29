// applied2 — a concurrency-safe store
// Integrate several ideas at once: a map for storage, a sync.Mutex for safety,
// methods for the API, and a sentinel error for the missing-key case.
// Implement Set.

// I AM NOT DONE
package main_test

import (
	"errors"
	"sync"
	"testing"
)

// ErrNotFound is returned by Get when a key is absent.
var ErrNotFound = errors.New("key not found")

// Store is a concurrency-safe string->int map.
type Store struct {
	mu sync.Mutex
	m  map[string]int
}

func NewStore() *Store {
	return &Store{m: make(map[string]int)}
}

// Set stores v under key, safely for concurrent callers.
func (s *Store) Set(key string, v int) {
	// FIXME: lock the mutex (defer unlock), then store the value.
}

// Get returns the value for key, or ErrNotFound if it is absent.
func (s *Store) Get(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.m[key]
	if !ok {
		return 0, ErrNotFound
	}
	return v, nil
}

func TestStoreConcurrent(t *testing.T) {
	s := NewStore()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Set("k", i)
		}(i)
	}
	wg.Wait()

	if _, err := s.Get("k"); err != nil {
		t.Errorf("expected k to be set, got %v", err)
	}
	if _, err := s.Get("missing"); !errors.Is(err, ErrNotFound) {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}
