// errors4
// panic/recover is Go's last-resort mechanism, not normal error handling.
// Use a deferred function to recover from a panic and turn it into an error.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// safeRun calls fn and, if fn panics, recovers and returns a non-nil error
// via the named return value. If fn does not panic, it returns nil.
func safeRun(fn func()) (err error) {
	// FIXME: add `defer func() { ... }()` that calls recover() and, when the
	// recovered value is non-nil, assigns an error to `err`
	// (e.g. fmt.Errorf("recovered: %v", r) — import "fmt").
	fn()
	return nil
}

func TestSafeRunRecovers(t *testing.T) {
	err := safeRun(func() { panic("boom") })
	if err == nil {
		t.Errorf("expected an error after recovering from a panic")
	}
}

func TestSafeRunNoPanic(t *testing.T) {
	err := safeRun(func() {})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
