// di2
// Inject behaviour through a small interface to make time-dependent code testable.

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// Clock reports the current time. Real code uses the system clock; a test
// injects a fixed one.
type Clock interface {
	Now() time.Time
}

// Greeting returns "Good morning" before noon, otherwise "Good evening",
// reading the time from the injected clock.
func Greeting(c Clock) string {
	// FIXME: if c.Now().Hour() is less than 12 return "Good morning",
	// otherwise return "Good evening".
	return ""
}

// fixedClock is a test double whose Now() always returns the same time.
type fixedClock struct{ t time.Time }

func (f fixedClock) Now() time.Time { return f.t }

func TestGreetingMorning(t *testing.T) {
	c := fixedClock{t: time.Date(2020, 1, 1, 9, 0, 0, 0, time.UTC)}
	if got := Greeting(c); got != "Good morning" {
		t.Errorf("want \"Good morning\", got %q", got)
	}
}

func TestGreetingEvening(t *testing.T) {
	c := fixedClock{t: time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC)}
	if got := Greeting(c); got != "Good evening" {
		t.Errorf("want \"Good evening\", got %q", got)
	}
}
