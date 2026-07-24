// structs1
// Define the struct's fields.
//
// A struct groups named fields, each with its own type.

package main_test

import "testing"

type Person struct {
	name string
	age  int
}

func TestPerson(t *testing.T) {
	person := Person{name: "John", age: 42}
	if person.name != "John" || person.age != 42 {
		t.Errorf("want name=John age=42, got name=%q age=%d", person.name, person.age)
	}
}
