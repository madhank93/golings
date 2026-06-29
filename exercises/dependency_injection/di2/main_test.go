// di2
// Inject behaviour through a small interface to make time-dependent code testable.

package main_test

import (
	"testing"
	"time"
)

type Clock interface {
	Now() time.Time
}

func Greeting(c Clock) string {
	if c.Now().Hour() < 12 {
		return "Good morning"
	}
	return "Good evening"
}

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
