// errors4
// Use a deferred function to recover from a panic and turn it into an error.

package main_test

import (
	"fmt"
	"testing"
)

func safeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	fn()
	return nil
}

func TestSafeRunRecovers(t *testing.T) {
	err := safeRun(func() { panic("boom") })
	if err == nil {
		t.Errorf("expected an error after recovering from a panic")
	}
}

func TestSafeRunOk(t *testing.T) {
	if err := safeRun(func() {}); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
