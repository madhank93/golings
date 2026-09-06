// process1
// Run a child process and capture what it printed.

package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// RunAndCapture runs a command and returns its standard output with trailing
// whitespace removed. Standard error is not part of the result.
//
// Output runs the command to completion and collects stdout for you; there is
// no need to wire up pipes and no need to call Wait. Reading from a pipe
// yourself and forgetting Wait leaves a zombie child behind.
func RunAndCapture(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(out), "\n"), nil
}

// The tests need a real child process, and the most portable one available is
// this test binary itself: re-run it with an environment marker, and the
// marker makes it behave as the child. It is the pattern os/exec's own tests
// use — no shell, no coreutils, works wherever `go test` works.
const helperEnv = "GOLINGS_PROCESS1_HELPER"

func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		t.Skip("not the child")
	}
	fmt.Fprintln(os.Stdout, "hello from the child")
	fmt.Fprintln(os.Stderr, "a warning nobody asked for")
	os.Exit(0)
}

// helper returns the command that re-runs this binary as the child.
func helper() (string, []string) {
	return os.Args[0], []string{"-test.run=^TestHelperProcess$"}
}

func withHelperEnv(t *testing.T) {
	t.Helper()
	t.Setenv(helperEnv, "1")
}

func TestRunAndCaptureStdout(t *testing.T) {
	withHelperEnv(t)

	name, args := helper()
	got, err := RunAndCapture(name, args...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "hello from the child") {
		t.Errorf("got %q, want it to contain %q", got, "hello from the child")
	}
}

func TestRunAndCaptureLeavesOutStderr(t *testing.T) {
	withHelperEnv(t)

	name, args := helper()
	got, err := RunAndCapture(name, args...)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "a warning nobody asked for") {
		t.Error("stderr leaked into the result — Output and CombinedOutput are different calls")
	}
}

func TestRunAndCaptureMissingCommand(t *testing.T) {
	if _, err := RunAndCapture("golings-no-such-command"); err == nil {
		t.Error("a command that does not exist must be an error")
	}
}
