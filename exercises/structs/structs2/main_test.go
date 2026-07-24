// structs2
// Embed a struct instead of adding a plain field.
//
// Embedding promotes the inner struct's fields onto the outer one, so they can
// be read directly — person.phone rather than person.contact.phone.

// I AM NOT DONE
package main_test

import "testing"

// FIXME: create a ContactDetails struct with a phone (string) field and EMBED
// it in Person — don't add a phone field directly to Person.
type Person struct {
	name string
	age  int
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
