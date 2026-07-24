// primitive_types1
// Toggle a boolean before the second check.
//
// A bool can be reassigned a new value; here the store opens, then closes, so
// each check fires exactly once.

package main_test

import (
	"reflect"
	"testing"
)

// storeMessages returns the message for each check the store passes.
func storeMessages() []string {
	var msgs []string

	storeIsOpen := true
	if storeIsOpen {
		msgs = append(msgs, "The store is open, let's buy some clothes!")
	}

	storeIsOpen = false

	if !storeIsOpen {
		msgs = append(msgs, "Oh no, let's buy some clothes online!")
	}

	return msgs
}

func TestStoreMessages(t *testing.T) {
	want := []string{
		"The store is open, let's buy some clothes!",
		"Oh no, let's buy some clothes online!",
	}
	if got := storeMessages(); !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
