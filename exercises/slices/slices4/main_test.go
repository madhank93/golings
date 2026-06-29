// slices4
// Index a slice within its bounds.

package main_test

import (
	"reflect"
	"testing"
)

func TestGetOnlyFirstName(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstName := names[0]
	if firstName != "John" {
		t.Errorf("firstName value must be John")
	}
}

func TestGetFirstTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstTwoNames := names[0:2]
	expected := []string{"John", "Maria"}
	if !reflect.DeepEqual(firstTwoNames, expected) {
		t.Errorf("firstTwoNames should be %v, but got %v", expected, firstTwoNames)
	}
}

func TestGetLastTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	lastTwoNames := names[2:4]
	expected := []string{"Carl", "Peter"}
	if !reflect.DeepEqual(lastTwoNames, expected) {
		t.Errorf("lastTwoNames should be %v, but got %v", expected, lastTwoNames)
	}
}
