// mock2
// A stub/fake returns canned data so you can test logic, including error paths.

package main_test

import (
	"errors"
	"testing"
)

type Fetcher interface {
	Fetch(id int) (string, error)
}

func WelcomeMessage(f Fetcher, id int) string {
	name, err := f.Fetch(id)
	if err != nil {
		return "Welcome, guest!"
	}
	return "Welcome, " + name + "!"
}

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
