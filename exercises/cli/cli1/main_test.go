// cli1
// Parse command-line flags with the flag package.

// I AM NOT DONE
package main_test

import (
	"flag"
	"testing"
)

// parseGreeting parses args like ["-name", "Go", "-count", "3"] and returns the
// values of the name and count flags. Using a FlagSet (rather than the global
// flag funcs) keeps it testable.
func parseGreeting(args []string) (string, int) {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	// FIXME: define the name and count flags, parse the args, and return their values.
	_ = fs
	return "", 0
}

func TestParseGreeting(t *testing.T) {
	name, count := parseGreeting([]string{"-name", "Go", "-count", "3"})
	if name != "Go" || count != 3 {
		t.Errorf("want (Go, 3), got (%q, %d)", name, count)
	}
}

func TestParseGreetingDefaults(t *testing.T) {
	name, count := parseGreeting(nil)
	if name != "world" || count != 1 {
		t.Errorf("want (world, 1), got (%q, %d)", name, count)
	}
}
