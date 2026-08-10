package logingest

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// record decodes the single JSON log line the middleware is expected to emit.
func record(t *testing.T, buf *bytes.Buffer) map[string]any {
	t.Helper()
	line := strings.TrimSpace(buf.String())
	if line == "" {
		t.Fatal("nothing was logged: is the handler actually wrapped?")
	}
	if strings.Count(line, "\n") > 0 {
		t.Fatalf("expected exactly one log record, got:\n%s", line)
	}
	var rec map[string]any
	if err := json.Unmarshal([]byte(line), &rec); err != nil {
		t.Fatalf("log line is not JSON (%v): %s", err, line)
	}
	return rec
}

// serve runs one request through the logging middleware and returns the
// recorder plus the captured log output.
func serve(t *testing.T, h http.Handler, method, target string) (*httptest.ResponseRecorder, *bytes.Buffer) {
	t.Helper()
	var buf bytes.Buffer
	wrapped := WithLogging(h, NewLogger(&buf, slog.LevelInfo))

	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, httptest.NewRequest(method, target, nil))
	return rec, &buf
}

func TestLoggerWritesJSONToTheGivenWriter(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(&buf, slog.LevelInfo)
	log.Info("hello", "key", "value")

	if buf.Len() == 0 {
		t.Fatal("NewLogger wrote nothing to the provided writer")
	}
	var rec map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &rec); err != nil {
		t.Fatalf("output is not JSON (%v): %s", err, buf.String())
	}
	if rec["msg"] != "hello" || rec["key"] != "value" {
		t.Errorf("record = %v, want msg=hello key=value", rec)
	}
}

func TestLoggerRespectsLevel(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(&buf, slog.LevelWarn)

	log.Info("should be dropped")
	if buf.Len() != 0 {
		t.Errorf("Info was logged at LevelWarn: %s", buf.String())
	}

	// Asserted so the test cannot pass simply because the logger writes
	// nowhere near buf.
	log.Warn("should be kept")
	if buf.Len() == 0 {
		t.Error("Warn was dropped at LevelWarn, or the logger is not writing to the given writer")
	}
}

func TestMiddlewareLogsMethodPathAndStatus(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})
	rec, buf := serve(t, h, http.MethodPost, "/events")

	if rec.Code != http.StatusCreated {
		t.Errorf("client saw status %d, want %d — the wrapper swallowed WriteHeader",
			rec.Code, http.StatusCreated)
	}

	got := record(t, buf)
	if got["msg"] != "request" {
		t.Errorf("msg = %v, want %q", got["msg"], "request")
	}
	if got["method"] != http.MethodPost {
		t.Errorf("method = %v, want %q", got["method"], http.MethodPost)
	}
	if got["path"] != "/events" {
		t.Errorf("path = %v, want %q", got["path"], "/events")
	}
	// JSON numbers decode as float64.
	if got["status"] != float64(http.StatusCreated) {
		t.Errorf("status = %v, want %d", got["status"], http.StatusCreated)
	}
	if _, ok := got["duration"]; !ok {
		t.Error("no duration attribute was logged")
	}
}

// A handler that writes a body without calling WriteHeader still sends 200.
// Recording 0 here is the bug this pins down.
func TestMiddlewareDefaultsStatusTo200(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte("ok")); err != nil {
			t.Errorf("write: %v", err)
		}
	})
	rec, buf := serve(t, h, http.MethodGet, "/healthz")

	if rec.Code != http.StatusOK {
		t.Errorf("client saw status %d, want %d", rec.Code, http.StatusOK)
	}
	if got := record(t, buf); got["status"] != float64(http.StatusOK) {
		t.Errorf("status = %v, want %d — a handler that never calls WriteHeader still sends 200",
			got["status"], http.StatusOK)
	}
}

// The middleware must be transparent: the body still reaches the client.
func TestMiddlewarePassesTheBodyThrough(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte("payload")); err != nil {
			t.Errorf("write: %v", err)
		}
	})
	rec, _ := serve(t, h, http.MethodGet, "/")

	if got := rec.Body.String(); got != "payload" {
		t.Errorf("body = %q, want %q", got, "payload")
	}
}

// One request, one record — a middleware that logs on the way in and again on
// the way out is twice the log volume for no extra information.
func TestMiddlewareLogsExactlyOneRecordPerRequest(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {})
	_, buf := serve(t, h, http.MethodGet, "/")

	// record fatals when the buffer is empty, so "logged nothing at all"
	// cannot pass this test by having no newlines in it.
	record(t, buf)
}
