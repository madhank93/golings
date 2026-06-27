// mock2
// A stub/fake returns canned data so you can test logic, including error paths.

// I AM NOT DONE
package main_test

import (
	"errors"
	"testing"
)

// Fetcher loads a user's name by id. The real one might hit a database.
type Fetcher interface {
	Fetch(id int) (string, error)
}

// WelcomeMessage returns "Welcome, <name>!" for a known user, or
// "Welcome, guest!" when Fetch returns an error.
func WelcomeMessage(f Fetcher, id int) string {
	// FIXME: call f.Fetch(id). If it returns an error, return "Welcome, guest!".
	// Otherwise return "Welcome, " + name + "!".
	return ""
}

// fakeFetcher returns canned results: a name for known ids, an error otherwise.
type fakeFetcher struct{ names map[int]string }

func (f fakeFetcher) Fetch(id int) (string, error) {
	if name, ok := f.names[id]; ok {
		return name, nil
	}
	return "", errors.New("not found")
}

func TestWelcomeKnownUser(t *testing.T) {
	f := fakeFetcher{names: map[int]string{1: "Go"}}
	if got := WelcomeMessage(f, 1); got != "Welcome, Go!" {
		t.Errorf("want %q, got %q", "Welcome, Go!", got)
	}
}

func TestWelcomeUnknownUser(t *testing.T) {
	f := fakeFetcher{names: map[int]string{}}
	if got := WelcomeMessage(f, 99); got != "Welcome, guest!" {
		t.Errorf("want %q, got %q", "Welcome, guest!", got)
	}
}
