// process3
// Tell a failed command apart from a command that never ran.

package main_test

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

// ExitStatus runs a command and reports the status it exited with.
//
// Two failures look alike at the call site and are not alike at all. A command
// that ran and exited non-zero is a *result*: grep says 1 for "no match", a
// test runner says 1 for "tests failed", and a shell reports that number
// happily. A command that could not be started — not on PATH, not executable —
// is an error in the caller's own world.
//
// exec.ExitError is how os/exec draws the line: err is an *ExitError exactly
// when the child ran. errors.As is the way to ask, because err may be wrapped.
func ExitStatus(name string, args ...string) (int, error) {
	err := exec.Command(name, args...).Run()
	if err == nil {
		return 0, nil
	}
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		return ee.ExitCode(), nil
	}
	return -1, err
}

const helperEnv = "GOLINGS_PROCESS3_HELPER"

// The child exits with whatever code the environment names.
func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) == "" {
		t.Skip("not the child")
	}
	code, err := strconv.Atoi(os.Getenv(helperEnv))
	if err != nil {
		code = 1
	}
	os.Exit(code)
}

func helper() (string, []string) {
	return os.Args[0], []string{"-test.run=^TestHelperProcess$"}
}

func TestExitStatusSuccess(t *testing.T) {
	t.Setenv(helperEnv, "0")

	name, args := helper()
	code, err := ExitStatus(name, args...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != 0 {
		t.Errorf("code = %d, want 0", code)
	}
}

func TestExitStatusNonZeroIsNotAnError(t *testing.T) {
	t.Setenv(helperEnv, "3")

	name, args := helper()
	code, err := ExitStatus(name, args...)
	if err != nil {
		t.Fatalf("a command that ran and exited 3 is a result, not an error: %v", err)
	}
	if code != 3 {
		t.Errorf("code = %d, want 3", code)
	}
}

func TestExitStatusCommandNotFound(t *testing.T) {
	code, err := ExitStatus("golings-no-such-command")
	if err == nil {
		t.Fatal("a command that never started must return an error")
	}
	if code != -1 {
		t.Errorf("code = %d, want -1", code)
	}
	if !errors.Is(err, exec.ErrNotFound) {
		t.Errorf("err = %v, want it to wrap exec.ErrNotFound", err)
	}
}
