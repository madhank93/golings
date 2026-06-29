// context3
// Contexts can carry request-scoped values down a call chain.

package main_test

import (
	"context"
	"testing"
)

type ctxKey string

const userKey ctxKey = "user"

func userFromContext(ctx context.Context) string {
	if u, ok := ctx.Value(userKey).(string); ok {
		return u
	}
	return "anonymous"
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
