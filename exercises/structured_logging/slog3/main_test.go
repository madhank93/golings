// slog3
// A slog.Handler decides what happens to each record. Implementing one lets you
// capture logs in memory for tests. Every method is provided except Handle,
// which must record the message.

// I AM NOT DONE
package main_test

import (
	"context"
	"log/slog"
	"testing"
)

// MemHandler collects the message of every record it handles.
type MemHandler struct {
	msgs *[]string
}

func (h MemHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h MemHandler) WithAttrs([]slog.Attr) slog.Handler       { return h }
func (h MemHandler) WithGroup(string) slog.Handler            { return h }

func (h MemHandler) Handle(_ context.Context, r slog.Record) error {
	// FIXME: append the record's message (r.Message) to the slice pointed to by
	// h.msgs. Right now it records nothing, so the test sees no logs.
	return nil
}

func TestMemHandler(t *testing.T) {
	var msgs []string
	logger := slog.New(MemHandler{msgs: &msgs})
	logger.Info("hello")
	logger.Warn("careful")
	if len(msgs) != 2 || msgs[0] != "hello" || msgs[1] != "careful" {
		t.Errorf("want [hello careful], got %v", msgs)
	}
}
