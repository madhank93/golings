// cli2
// Read the positional arguments left after flag parsing with FlagSet.Args.

// I AM NOT DONE
package main_test

import (
	"flag"
	"testing"
)

// firstArg parses flags from args and returns the first positional (non-flag)
// argument, or "" when there is none — e.g. ["-v", "build", "./..."] -> "build".
func firstArg(args []string) string {
	fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
	verbose := fs.Bool("v", false, "verbose output")
	_ = verbose
	// FIXME: parse the args, then return the first positional argument.
	return ""
}

func TestFirstArg(t *testing.T) {
	if got := firstArg([]string{"-v", "build", "./..."}); got != "build" {
		t.Errorf("want \"build\", got %q", got)
	}
}

func TestFirstArgNone(t *testing.T) {
	if got := firstArg([]string{"-v"}); got != "" {
		t.Errorf("want \"\", got %q", got)
	}
}
