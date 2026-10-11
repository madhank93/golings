package app_test

import (
	"testing"

	"example.com/app"
)

func TestWelcome(t *testing.T) {
	if got := app.Welcome("Ada"); got != "Good day, Ada." {
		t.Errorf("Welcome(%q) = %q", "Ada", got)
	}
}
