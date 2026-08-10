package logingest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// A fixed timestamp keeps the golden JSON stable; a real clock in a test is a
// flake waiting to happen.
var testAt = time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC)

func newTestEvent(source, msg string) Event {
	return Event{Source: source, Level: "info", Message: msg, At: testAt}
}

func mustJSON(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(b)
}

// do drives the handler in-process. httptest.NewRecorder is enough here — a
// real listener would add a port, a timeout and a source of flakiness without
// testing anything extra.
func do(t *testing.T, s *Server, method, target, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	return rec
}

// A well-formed event is accepted and reaches the store.
func TestIngestAcceptsValidEvent(t *testing.T) {
	store := NewStore()
	s := NewServer(store)

	rec := do(t, s, http.MethodPost, "/events", mustJSON(t, newTestEvent("api", "hello")))

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
	if store.Len() != 1 {
		t.Errorf("store.Len() = %d, want 1", store.Len())
	}
}

// A body that is not JSON is the client's mistake, not a server error.
func TestIngestRejectsMalformedJSON(t *testing.T) {
	store := NewStore()
	s := NewServer(store)

	rec := do(t, s, http.MethodPost, "/events", "{not json")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if store.Len() != 0 {
		t.Errorf("store.Len() = %d, want 0", store.Len())
	}
}

// Valid JSON can still be an invalid event. ErrInvalidEvent must surface as 400.
func TestIngestRejectsInvalidEvent(t *testing.T) {
	store := NewStore()
	s := NewServer(store)

	rec := do(t, s, http.MethodPost, "/events", mustJSON(t, newTestEvent("", "no source")))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if store.Len() != 0 {
		t.Errorf("store.Len() = %d, want 0", store.Len())
	}
}

// The {source} wildcard has to reach the handler, and order must survive.
func TestBySourceReturnsEventsInOrder(t *testing.T) {
	store := NewStore()
	for _, msg := range []string{"first", "second"} {
		if err := store.Add(newTestEvent("api", msg)); err != nil {
			t.Fatalf("seed: %v", err)
		}
	}
	s := NewServer(store)

	rec := do(t, s, http.MethodGet, "/events/api", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got []Event
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode body %q: %v", rec.Body.String(), err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d events, want 2", len(got))
	}
	if got[0].Message != "first" || got[1].Message != "second" {
		t.Errorf("got %q then %q, want %q then %q",
			got[0].Message, got[1].Message, "first", "second")
	}
}

// The raw body matters, not just the decoded value: a JavaScript client doing
// data.forEach() breaks on null and works on [].
func TestBySourceUnknownEncodesEmptyArrayNotNull(t *testing.T) {
	s := NewServer(NewStore())

	rec := do(t, s, http.MethodGet, "/events/nobody", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "[]" {
		t.Errorf("body = %q, want %q", got, "[]")
	}
}

// This is what the method pattern buys: the mux rejects the verb itself, so no
// handler code runs and you never write the check by hand.
func TestUnsupportedMethodIsRejectedByTheMux(t *testing.T) {
	s := NewServer(NewStore())

	rec := do(t, s, http.MethodDelete, "/events", "")

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
