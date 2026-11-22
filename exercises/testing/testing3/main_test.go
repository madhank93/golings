// testing3
// Make the test pass!
//
// greet is correct. TestGreetEmpty checks the failure case, but it was copied
// from a happy-path test and expects the opposite of what should happen.

package main_test

import (
	"errors"
	"testing"
)

var ErrEmptyName = errors.New("empty name")

func greet(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + name, nil
}

func TestGreet(t *testing.T) {
	got, err := greet("Ana")
	if err != nil {
		t.Fatalf("greet(%q) error: %v", "Ana", err)
	}
	if want := "Hello, Ana"; got != want {
		t.Errorf("greet(%q) = %q, want %q", "Ana", got, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	_, err := greet("")
	if !errors.Is(err, ErrEmptyName) {
		t.Fatalf("greet(%q) error = %v, want ErrEmptyName", "", err)
	}
}
