// process2
// Feed a child process on stdin and read what it writes back.

// I AM NOT DONE
package main_test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Filter runs a command with input as its standard input and returns its
// standard output, trailing newline removed.
//
// Setting Stdin to a reader lets os/exec do the plumbing: it copies the reader
// into the pipe and — the part that matters — closes the pipe when the reader
// is exhausted. A child reading until EOF never sees one otherwise, and both
// processes wait for each other forever.
func Filter(input, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	// FIXME: give the child `input` on its standard input, run it, and return
	// its stdout with the trailing newline trimmed. The child reads until EOF,
	// so whatever you attach has to end — a pipe you never close is a deadlock,
	// not an error.
	_ = strings.NewReader
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

const helperEnv = "GOLINGS_PROCESS2_HELPER"

// The child upper-cases every line it is given, until stdin reaches EOF.
func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		t.Skip("not the child")
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println(strings.ToUpper(sc.Text()))
	}
	os.Exit(0)
}

func helper() (string, []string) {
	return os.Args[0], []string{"-test.run=^TestHelperProcess$"}
}

func TestFilter(t *testing.T) {
	t.Setenv(helperEnv, "1")

	name, args := helper()
	got, err := Filter("one\ntwo\n", name, args...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "ONE\nTWO" {
		t.Errorf("got %q, want %q", got, "ONE\nTWO")
	}
}

// A child that reads to EOF only finishes when its stdin is closed. If this
// test hangs until the runner's timeout, that is the bug.
func TestFilterClosesStdin(t *testing.T) {
	t.Setenv(helperEnv, "1")

	name, args := helper()
	got, err := Filter("only one line\n", name, args...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "ONLY ONE LINE" {
		t.Errorf("got %q, want %q", got, "ONLY ONE LINE")
	}
}

func TestFilterEmptyInput(t *testing.T) {
	t.Setenv(helperEnv, "1")

	name, args := helper()
	got, err := Filter("", name, args...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Errorf("got %q, want empty", got)
	}
}
