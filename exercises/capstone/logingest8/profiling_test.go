package logingest

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var svcAt = time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC)

func svcEvent(source, msg string) Event {
	return Event{Source: source, Level: "info", Message: msg, At: svcAt}
}

func get(t *testing.T, h http.Handler, target string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, target, nil))
	return rec
}

func TestWithPprofServesTheIndex(t *testing.T) {
	mux := WithPprof(http.NewServeMux())

	rec := get(t, mux, PprofPrefix)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d — are the pprof handlers registered?", rec.Code, http.StatusOK)
	}
	// The index lists the named profiles; heap is always one of them.
	if !strings.Contains(rec.Body.String(), "heap") {
		t.Error("the pprof index did not list the heap profile")
	}
}

// pprof.Index also serves the named profiles, so registering it on the prefix
// is enough to reach /debug/pprof/heap.
func TestWithPprofServesNamedProfiles(t *testing.T) {
	mux := WithPprof(http.NewServeMux())

	rec := get(t, mux, PprofPrefix+"heap?debug=1")

	if rec.Code != http.StatusOK {
		t.Errorf("heap profile status = %d, want %d", rec.Code, http.StatusOK)
	}
	if rec.Body.Len() == 0 {
		t.Error("heap profile body was empty")
	}
}

// Only the profiles are exposed — registering the handlers explicitly must not
// drag in unrelated routes.
func TestWithPprofDoesNotExposeOtherPaths(t *testing.T) {
	mux := WithPprof(http.NewServeMux())

	if rec := get(t, mux, "/"); rec.Code != http.StatusNotFound {
		t.Errorf("GET / status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestServiceStillServesTheEventRoutes(t *testing.T) {
	store := NewStore()
	if err := store.Add(svcEvent("api", "seeded")); err != nil {
		t.Fatalf("seed: %v", err)
	}
	var logBuf bytes.Buffer
	svc := NewService(store, NewLogger(&logBuf, slog.LevelInfo))

	rec := get(t, svc, "/events/api")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got []Event
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode %q: %v", rec.Body.String(), err)
	}
	if len(got) != 1 || got[0].Message != "seeded" {
		t.Errorf("got %+v, want one event with message %q", got, "seeded")
	}
}

// The pprof prefix is longer than the "/" catch-all, so ServeMux must prefer it.
func TestServiceServesProfilesAlongsideEvents(t *testing.T) {
	var logBuf bytes.Buffer
	svc := NewService(NewStore(), NewLogger(&logBuf, slog.LevelInfo))

	rec := get(t, svc, PprofPrefix)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d — the catch-all route is shadowing the profiles",
			rec.Code, http.StatusOK)
	}
}

// The logger is outermost, so everything that goes through the service is
// logged — profile requests included.
func TestServiceLogsEveryRequest(t *testing.T) {
	var logBuf bytes.Buffer
	svc := NewService(NewStore(), NewLogger(&logBuf, slog.LevelInfo))

	get(t, svc, "/events/api")

	line := strings.TrimSpace(logBuf.String())
	if line == "" {
		t.Fatal("nothing was logged: is WithLogging wrapped around the mux?")
	}
	var rec map[string]any
	if err := json.Unmarshal([]byte(line), &rec); err != nil {
		t.Fatalf("log line is not JSON (%v): %s", err, line)
	}
	if rec["path"] != "/events/api" {
		t.Errorf("logged path = %v, want %q", rec["path"], "/events/api")
	}
	if rec["status"] != float64(http.StatusOK) {
		t.Errorf("logged status = %v, want %d", rec["status"], http.StatusOK)
	}
}

// The full path an event takes through the finished service: submitted through
// the pool, stored, then read back over HTTP.
func TestServiceEndToEndThroughThePool(t *testing.T) {
	store := NewStore()
	var logBuf bytes.Buffer
	svc := NewService(store, NewLogger(&logBuf, slog.LevelInfo))

	pool := NewPool(store, 4)
	for i := range 10 {
		pool.Submit(svcEvent("worker", "batched "+string(rune('0'+i))))
	}
	pool.Close()

	rec := get(t, svc, "/events/worker")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got []Event
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode %q: %v", rec.Body.String(), err)
	}
	if len(got) != 10 {
		t.Errorf("got %d events, want 10", len(got))
	}
	if pool.Rejected() != 0 {
		t.Errorf("Rejected() = %d, want 0", pool.Rejected())
	}
}
