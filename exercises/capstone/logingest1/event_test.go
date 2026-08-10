package logingest

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func validEvent() Event {
	return Event{
		Source:  "api",
		Level:   "info",
		Message: "request served",
		At:      time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC),
	}
}

func TestValidateAcceptsWellFormedEvent(t *testing.T) {
	for _, level := range Levels {
		e := validEvent()
		e.Level = level
		if err := e.Validate(); err != nil {
			t.Errorf("level %q: Validate() = %v, want nil", level, err)
		}
	}
}

func TestValidateRejectsMalformedEvents(t *testing.T) {
	// field is the word the error text must mention, so a caller reading the
	// message can tell which part of the payload was wrong.
	cases := []struct {
		name   string
		field  string
		mutate func(*Event)
	}{
		{"empty source", "source", func(e *Event) { e.Source = "" }},
		{"unknown level", "level", func(e *Event) { e.Level = "trace" }},
		{"empty level", "level", func(e *Event) { e.Level = "" }},
		{"empty message", "message", func(e *Event) { e.Message = "" }},
		{"zero timestamp", "at", func(e *Event) { e.At = time.Time{} }},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := validEvent()
			tc.mutate(&e)

			err := e.Validate()
			if err == nil {
				t.Fatal("Validate() = nil, want an error")
			}
			// The sentinel is the machine-readable half of the contract.
			if !errors.Is(err, ErrInvalidEvent) {
				t.Errorf("errors.Is(%v, ErrInvalidEvent) = false; wrap it with %%w", err)
			}
			// The wrapped text is the human-readable half.
			if !strings.Contains(strings.ToLower(err.Error()), tc.field) {
				t.Errorf("error %q does not mention %q", err, tc.field)
			}
		})
	}
}
