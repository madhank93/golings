## logingest2 — an RWMutex, and why you return a copy

```go
func (s *Store) Add(e Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[e.Source] = append(s.events[e.Source], e)
	return nil
}

func (s *Store) BySource(source string) []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if got := s.events[source]; got != nil {
		return slices.Clone(got)
	}
	return []Event{}
}

func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	for _, evs := range s.events {
		n += len(evs)
	}
	return n
}
```

**Why it works**

- `sync.RWMutex` splits the lock in two. `Lock` is exclusive for writers;
  `RLock` lets any number of readers in at once. A store that is read far more
  than it is written — which is most stores — gets real concurrency out of that
  split, where a plain `sync.Mutex` would serialise every reader.
- `defer s.mu.Unlock()` on the line after `Lock` is the habit worth building.
  An early `return` added six months later cannot forget to unlock.
- Validation happens **before** the lock. Rejecting a bad event never needs the
  map, and holding a lock across work that does not need it is how contention
  starts.
- Returning the validation error unchanged, rather than wrapping it again, is
  what keeps `errors.Is(err, ErrInvalidEvent)` working for the HTTP handler in
  stage three.

**Key detail:** `slices.Clone` is not decoration. `s.events[source]` is a slice
header pointing at the store's own backing array — hand it out and the caller
can write straight into your map's data with no lock held, and a later `append`
by the store can overwrite what the caller is still reading. The lock protects
the map only while you hold it; the slice escapes it. Copy on the way out.

**Key detail:** `BySource` returns `[]Event{}` rather than the nil slice for an
unknown source. Both have length zero and both range fine, but stage three
marshals this to JSON — and `nil` becomes `null` while an empty slice becomes
`[]`. A client that does `data.forEach(...)` breaks on one and not the other.

**Key detail:** concurrent *map writes* are not a subtle race the detector might
miss. The Go runtime checks for them directly and kills the process with
`fatal error: concurrent map writes` — unrecoverable, no panic to defer against.
Run the concurrency test with `-race` and you will see it.

**References**

- pkg.go.dev — sync.RWMutex: https://pkg.go.dev/sync#RWMutex
- pkg.go.dev — slices.Clone: https://pkg.go.dev/slices#Clone
- Go blog — Introducing the Go Race Detector: https://go.dev/blog/race-detector
