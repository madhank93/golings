// errors3
// A custom error type is any type implementing the error interface.
// errors.As extracts a wrapped error of a concrete type.
// Implement ValidationError.Error so the message is built from its field.

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

// ValidationError is a custom error type carrying the offending field.
type ValidationError struct {
	Field string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	// FIXME: build the message from the field with fmt.
	return ""
}

func validate(name string) error {
	if name == "" {
		return fmt.Errorf("validate: %w", &ValidationError{Field: "name"})
	}
	return nil
}

func TestValidationErrorAs(t *testing.T) {
	err := validate("")

	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected a *ValidationError, got %v", err)
	}
	if ve.Field != "name" {
		t.Errorf("want field %q, got %q", "name", ve.Field)
	}
	if ve.Error() != "name is required" {
		t.Errorf("want message %q, got %q", "name is required", ve.Error())
	}
}
