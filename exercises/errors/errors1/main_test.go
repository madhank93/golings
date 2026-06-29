// errors1
// In Go, errors are ordinary values returned alongside results.

package main_test

import (
	"errors"
	"testing"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
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
