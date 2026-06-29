// errors2
// Wrapping an error with %w keeps the original reachable via errors.Is.

package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotFound = errors.New("not found")

func lookup(id string) error {
	if id == "" {
		return fmt.Errorf("lookup %q: %w", id, ErrNotFound)
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
