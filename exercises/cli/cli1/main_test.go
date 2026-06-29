// cli1
// Parse command-line flags with the flag package.

package main_test

import (
	"flag"
	"testing"
)

func parseGreeting(args []string) (string, int) {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	name := fs.String("name", "world", "")
	count := fs.Int("count", 1, "")
	_ = fs.Parse(args)
	return *name, *count
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
