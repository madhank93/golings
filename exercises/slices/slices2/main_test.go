// slices2
// Take a sub-slice using [low:high] bounds.
//
// A sub-slice includes low up to but NOT including high, so the last two of a
// length-4 array are [2:4].

// I AM NOT DONE
package main_test

import (
	"reflect"
	"testing"
)

// FIXME: return the last two names using [low:high] bounds.
func lastTwo(names [4]string) []string {
	return names[]
}

func TestLastTwo(t *testing.T) {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
	want := []string{"Carl", "Peter"}
	if got := lastTwo(names); !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
