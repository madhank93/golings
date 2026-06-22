// context3
// Contexts can carry request-scoped values down a call chain.
// Read a value out of the context, falling back to a default when absent.

// I AM NOT DONE
package main_test

import (
	"context"
	"testing"
)

// ctxKey is an unexported type for context keys to avoid collisions.
type ctxKey string

const userKey ctxKey = "user"

// userFromContext returns the user stored under userKey, or "anonymous" if none.
func userFromContext(ctx context.Context) string {
	// FIXME: ctx.Value(userKey) returns an interface{}. Type-assert it to
	// string with the comma-ok form; return it when present, else "anonymous".
	return ""
}

func TestUserPresent(t *testing.T) {
	ctx := context.WithValue(context.Background(), userKey, "alice")
	if got := userFromContext(ctx); got != "alice" {
		t.Errorf("want alice, got %q", got)
	}
}

func TestUserAbsent(t *testing.T) {
	if got := userFromContext(context.Background()); got != "anonymous" {
		t.Errorf("want anonymous, got %q", got)
	}
}
