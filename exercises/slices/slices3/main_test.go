// slices3
// Add elements to a slice with append.
//
// append returns a new slice with the extra elements added to the end.

// I AM NOT DONE
package main_test

import (
	"reflect"
	"testing"
)

// FIXME: append "Peter" to names and return the result.
func addPeter(names []string) []string {
	return append()
}

func TestAddPeter(t *testing.T) {
	names := []string{"John", "Maria", "Carl"}
	want := []string{"John", "Maria", "Carl", "Peter"}
	if got := addPeter(names); !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
