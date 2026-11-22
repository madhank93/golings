// testing1
// Make the test run!
//
// `go test` only runs functions shaped exactly like TestXxx(t *testing.T).
// This one has the right name but the wrong parameter type.

package main_test

import "testing"

func double(n int) int {
	return n * 2
}

func TestDouble(t *testing.T) {
	if got := double(4); got != 8 {
		t.Errorf("double(4) = %d, want 8", got)
	}
}
