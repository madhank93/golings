// errors5
// errors.Join (Go 1.20) bundles several errors into one.

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

func validatePassword(p string) error {
	var errs []error
	if len(p) < 8 {
		errs = append(errs, ErrTooShort)
	}
	if !strings.ContainsAny(p, "0123456789") {
		errs = append(errs, ErrNoDigit)
	}
	return errors.Join(errs...)
}

func TestValidatePassword(t *testing.T) {
	err := validatePassword("abc")
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
