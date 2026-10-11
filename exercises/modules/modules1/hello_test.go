package hello_test

import (
	"testing"

	"example.com/hello"
)

func TestGreeting(t *testing.T) {
	if got := hello.Greeting("gopher"); got != "hello, gopher" {
		t.Errorf("Greeting(%q) = %q, want %q", "gopher", got, "hello, gopher")
	}
}
