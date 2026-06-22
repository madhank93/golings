// errors2
// Wrapping an error with %w keeps the original reachable via errors.Is.
// Make lookup wrap the ErrNotFound sentinel.

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

// ErrNotFound is a sentinel error callers can test for with errors.Is.
var ErrNotFound = errors.New("not found")

// lookup returns an error that WRAPS ErrNotFound when id is empty.
func lookup(id string) error {
	if id == "" {
		// FIXME: add context AND wrap the sentinel using the %w verb,
		// e.g. fmt.Errorf("lookup %q: %w", id, ErrNotFound)
		return fmt.Errorf("lookup failed")
	}
	return nil
}

func TestLookupWrapsSentinel(t *testing.T) {
	err := lookup("")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected the error to wrap ErrNotFound, got %v", err)
	}
}

func TestLookupOk(t *testing.T) {
	if err := lookup("abc"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
