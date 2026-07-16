// slog2
// logger.With(attrs...) returns a NEW logger that carries those attributes on
// every record it writes. It does not modify the receiver — you have to use the
// logger it returns.

package main_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

// logWithService logs msg and must tag every line with service=api.
func logWithService(w *bytes.Buffer, msg string) {
	logger := slog.New(slog.NewTextHandler(w, nil))
	// FIXME: With() returns a new logger; its result is discarded here, so the
	// service attribute never reaches the output. Assign it back to logger.
	logger = logger.With("service", "api")
	logger.Info(msg)
}

func TestLogWithService(t *testing.T) {
	var buf bytes.Buffer
	logWithService(&buf, "ok")
	if out := buf.String(); !strings.Contains(out, "service=api") {
		t.Errorf("expected service=api attribute, got %q", out)
	}
}
