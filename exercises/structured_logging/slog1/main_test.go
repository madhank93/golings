// slog1
// log/slog (Go 1.21) is structured logging in the stdlib. A handler has a
// minimum level and drops records below it. Here the handler is pinned to
// LevelWarn, so Info messages vanish.

package main_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

// logInfo writes an Info-level message to w using a text handler.
func logInfo(w *bytes.Buffer, msg string) {
	// FIXME: the handler drops anything below LevelWarn, so Info never appears.
	// Set Level to slog.LevelInfo.
	h := slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.New(h).Info(msg)
}

func TestLogInfo(t *testing.T) {
	var buf bytes.Buffer
	logInfo(&buf, "server started")
	if !strings.Contains(buf.String(), "server started") {
		t.Errorf("expected the Info line in output, got %q", buf.String())
	}
}
