// structs3
// Attach a method to a struct.
//
// A method returning a formatted string is easy to get subtly wrong — a missing
// space between the names still compiles. The test pins the exact result.

package main_test

import "testing"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

func TestFullName(t *testing.T) {
	person := Person{firstName: "Maurício", lastName: "Antunes"}
	if got := person.FullName(); got != "Maurício Antunes" {
		t.Errorf(`want "Maurício Antunes", got %q`, got)
	}
}
