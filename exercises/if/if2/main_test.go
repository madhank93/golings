// if2
// Map each input string to an output with if/else.

package main_test

import "testing"

func fooIfFizz(fizzish string) string {
	if fizzish == "fizz" {
		return "foo"
	} else if fizzish == "fuzz" {
		return "bar"
	}
	return "baz"
}

func TestFooForFizz(t *testing.T) {
	if result := fooIfFizz("fizz"); result != "foo" {
		t.Errorf("should be 'foo' but got %s", result)
	}
}

func TestBarForFuzz(t *testing.T) {
	if result := fooIfFizz("fuzz"); result != "bar" {
		t.Errorf("should be 'bar' but got %s", result)
	}
}

func TestDefaultForBazz(t *testing.T) {
	if result := fooIfFizz("random stuff"); result != "baz" {
		t.Errorf("should be 'baz' but got %s", result)
	}
}
