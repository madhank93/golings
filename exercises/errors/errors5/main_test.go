// errors5
// errors.Join (Go 1.20) bundles several errors into one. errors.Is then
// matches ANY of the joined errors — handy for collecting validation failures.

// I AM NOT DONE
package main_test

import (
	"errors"
	"strings"
	"testing"
)

var (
	ErrTooShort = errors.New("too short")
	ErrNoDigit  = errors.New("no digit")
)

// validatePassword returns a single error combining every rule that failed,
// or nil if the password is valid.
func validatePassword(p string) error {
	var errs []error
	if len(p) < 8 {
		errs = append(errs, ErrTooShort)
	}
	if !strings.ContainsAny(p, "0123456789") {
		errs = append(errs, ErrNoDigit)
	}
	// FIXME: combine errs into ONE error with errors.Join(errs...) and return it.
	// (errors.Join returns nil when errs is empty, so a valid password -> nil.)
	return nil
}

func TestValidatePassword(t *testing.T) {
	err := validatePassword("abc") // both rules fail
	if !errors.Is(err, ErrTooShort) {
		t.Error("expected joined error to contain ErrTooShort")
	}
	if !errors.Is(err, ErrNoDigit) {
		t.Error("expected joined error to contain ErrNoDigit")
	}
	if validatePassword("password1") != nil {
		t.Error("a valid password should return nil")
	}
}
