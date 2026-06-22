package exercises

import (
	"errors"
	"os/exec"
	"path/filepath"
)

// Status is the rich, gated state of an exercise used by the TUI. It is finer
// than State: an exercise only reaches StatusDone when its marker is removed,
// it compiles/tests cleanly, AND golangci-lint is clean.
type Status int

const (
	StatusPending  Status = iota // "I AM NOT DONE" marker still present
	StatusFailing                // marker removed, but Run() fails
	StatusLintFail               // marker removed, Run() passes, lint not clean
	StatusDone                   // marker removed, Run() passes, lint clean
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusFailing:
		return "Failing"
	case StatusLintFail:
		return "LintFail"
	case StatusDone:
		return "Done"
	default:
		return "Unknown"
	}
}

// Verify computes the gated status of e. It short-circuits on the marker (no
// run), then runs the exercise, then lints it only if the run passed.
// The returned Result carries the output to show the learner.
func (e Exercise) Verify() (Status, Result) {
	if e.State() == Pending {
		return StatusPending, Result{Exercise: e}
	}

	result, err := e.Run()
	if err != nil {
		return StatusFailing, result
	}

	if ok, out := Lint(e); !ok {
		result.Out = out
		result.Err = ""
		return StatusLintFail, result
	}

	return StatusDone, result
}

// Lint runs golangci-lint on the exercise's directory. It returns (true, "")
// when clean. If golangci-lint is not installed it degrades gracefully
// (returns true with a note) rather than blocking progress forever.
func Lint(e Exercise) (bool, string) {
	dir := filepath.Dir(e.Path)
	cmd := exec.Command("golangci-lint", "run", "./"+dir)
	out, err := cmd.CombinedOutput()
	if err == nil {
		return true, ""
	}

	var execErr *exec.Error
	if errors.As(err, &execErr) {
		// binary missing / not executable — skip the gate
		return true, "golangci-lint not found; lint gate skipped"
	}

	// non-zero exit => lint findings
	return false, string(out)
}
