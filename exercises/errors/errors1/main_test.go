// errors1
// In Go, errors are ordinary values returned alongside results.
// Make divide return an error when the divisor is zero.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// divide should return (quotient, nil) normally, and (0, error) when b == 0.
func divide(a, b int) (int, error) {
	if b == 0 {
		// FIXME: a zero divisor is invalid. Return 0 and an error
		// created with errors.New (you'll need to import "errors").
		return 0, nil
	}
	return a / b, nil
}

func TestDivideOk(t *testing.T) {
	got, err := divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("want 5, got %d", got)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(1, 0)
	if err == nil {
		t.Errorf("expected an error when dividing by zero")
	}
}
