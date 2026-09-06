package logingest

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func testEvent(source, msg string) Event {
	return Event{
		Source:  source,
		Level:   "info",
		Message: msg,
		At:      time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC),
	}
}

func TestAddRejectsInvalidEvent(t *testing.T) {
	s := NewStore()
	bad := testEvent("", "no source")

	err := s.Add(bad)
	if !errors.Is(err, ErrInvalidEvent) {
		t.Errorf("Add(invalid) = %v, want an error wrapping ErrInvalidEvent", err)
	}
	if s.Len() != 0 {
		t.Errorf("store kept a rejected event: Len() = %d, want 0", s.Len())
	}
}

func TestAddAndBySource(t *testing.T) {
	s := NewStore()
	for i := range 3 {
		if err := s.Add(testEvent("api", fmt.Sprintf("msg %d", i))); err != nil {
			t.Fatalf("Add: %v", err)
		}
	}
	if err := s.Add(testEvent("worker", "other")); err != nil {
		t.Fatalf("Add: %v", err)
	}

	got := s.BySource("api")
	if len(got) != 3 {
		t.Fatalf("BySource(\"api\") returned %d events, want 3", len(got))
	}
	for i, e := range got {
		if want := fmt.Sprintf("msg %d", i); e.Message != want {
			t.Errorf("event %d message = %q, want %q (order must be preserved)", i, e.Message, want)
		}
	}
	if s.Len() != 4 {
		t.Errorf("Len() = %d, want 4", s.Len())
	}
}

func TestBySourceUnknownReturnsEmptyNotNil(t *testing.T) {
	s := NewStore()
	got := s.BySource("nobody")
	if got == nil {
		t.Error("BySource(unknown) = nil, want an empty slice")
	}
	if len(got) != 0 {
		t.Errorf("BySource(unknown) returned %d events, want 0", len(got))
	}
}

// A caller must not be able to reach back into the store through the slice it
// was handed.
func TestBySourceReturnsACopy(t *testing.T) {
	s := NewStore()
	if err := s.Add(testEvent("api", "original")); err != nil {
		t.Fatalf("Add: %v", err)
	}

	got := s.BySource("api")
	if len(got) != 1 {
		t.Fatalf("BySource(\"api\") returned %d events, want 1", len(got))
	}
	got[0].Message = "tampered"

	again := s.BySource("api")
	if len(again) != 1 {
		t.Fatalf("BySource(\"api\") returned %d events on the second call, want 1", len(again))
	}
	if again[0].Message != "original" {
		t.Errorf("mutating the returned slice changed the store: got %q, want %q",
			again[0].Message, "original")
	}
}

// Run with -race. Without locking this fails as a detected data race, or as a
// "concurrent map writes" fatal error, long before the count is checked.
func TestStoreIsSafeForConcurrentUse(t *testing.T) {
	s := NewStore()

	const writers, each = 8, 50
	var wg sync.WaitGroup
	for w := range writers {
		wg.Go(func() {
			for i := range each {
				if err := s.Add(testEvent(fmt.Sprintf("src-%d", w), fmt.Sprintf("msg %d", i))); err != nil {
					t.Errorf("Add: %v", err)
					return
				}
			}
		})
	}
	// Readers run concurrently with the writers so the race detector sees a
	// read and a write overlap, not just two writes.
	for range writers {
		wg.Go(func() {
			for range each {
				_ = s.BySource("src-0")
				_ = s.Len()
			}
		})
	}
	wg.Wait()

	if got := s.Len(); got != writers*each {
		t.Errorf("Len() = %d, want %d", got, writers*each)
	}
}
