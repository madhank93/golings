// cli2
// Read the positional arguments left after flag parsing with FlagSet.Args.

package main_test

import (
	"flag"
	"testing"
)

func firstArg(args []string) string {
	fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
	_ = fs.Bool("v", false, "verbose output")
	_ = fs.Parse(args)
	return fs.Arg(0)
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
