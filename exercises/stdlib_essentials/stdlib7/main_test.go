// stdlib7 — encoding/json omitzero
// `omitempty` predates most of Go's type system: it drops empty strings, zero
// numbers, false, and nil/empty collections — but a struct is never "empty" to
// it, so a zero time.Time marshals anyway. Go 1.24 added `omitzero`, which asks
// the type itself whether it is the zero value.

package main_test

import (
	"encoding/json"
	"testing"
	"time"
)

// Event should omit At entirely when it was never set.
type Event struct {
	Name string    `json:"name"`
	At   time.Time `json:"at,omitzero"`
}

func encodeEvent(e Event) (string, error) {
	data, err := json.Marshal(e)
	return string(data), err
}

func TestEncodeEventUnset(t *testing.T) {
	got, err := encodeEvent(Event{Name: "launch"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := `{"name":"launch"}`; got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestEncodeEventSet(t *testing.T) {
	at := time.Date(2026, 3, 1, 12, 0, 0, 0, time.UTC)
	got, err := encodeEvent(Event{Name: "launch", At: at})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := `{"name":"launch","at":"2026-03-01T12:00:00Z"}`; got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
