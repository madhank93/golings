// structs2
// Embed a struct instead of adding a plain field.
//
// Embedding promotes the inner struct's fields onto the outer one, so they can
// be read directly — person.phone rather than person.contact.phone.

package main_test

import "testing"

type ContactDetails struct {
	phone string
}

type Person struct {
	name string
	age  int
	ContactDetails
}

func TestPhone(t *testing.T) {
	person := Person{
		name:           "John",
		age:            32,
		ContactDetails: ContactDetails{phone: "+01 101 102"},
	}
	if person.phone != "+01 101 102" {
		t.Errorf(`want phone "+01 101 102", got %q`, person.phone)
	}
}
