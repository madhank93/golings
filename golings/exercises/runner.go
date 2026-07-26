package exercises

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// DefaultTimeout bounds a single exercise run. An exercise that deadlocks or
// loops forever is a normal thing for a learner to write; without a bound it
// wedged the TUI with no way out but killing the terminal.
const DefaultTimeout = 60 * time.Second

// maxOutput caps how much of a run's stdout/stderr is retained. A runaway
// print loop emits gigabytes in seconds, and all of it was being held in
// memory and then handed to the renderer.
const maxOutput = 1 << 20 // 1 MiB per stream

type Result struct {
	Exercise Exercise
	Out      string
	Err      string
}

// Run executes the exercise under DefaultTimeout.
func (e Exercise) Run() (Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	return e.RunContext(ctx)
}

// RunContext executes the exercise, stopping it when ctx is done.
//
// The command runs in its own process group and cancellation signals the
// group, not just the direct child: `go test` and `go run` are supervisors
// that spawn the compiled binary, so killing the parent alone leaves the
// exercise itself running and holding its stdout pipe.
func (e Exercise) RunContext(ctx context.Context) (Result, error) {
	cmd := exec.CommandContext(ctx, "go", BuildArgs(e)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		// Negative pid addresses the whole group. SIGKILL rather than SIGTERM:
		// the point of a cancel is that it takes effect now.
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
	// A grandchild that inherited the stdout pipe can outlive the group kill;
	// without a bound, Wait blocks on pipe EOF forever.
	cmd.WaitDelay = 3 * time.Second

	var stdout, stderr cappedBuffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	res := Result{Exercise: e, Out: stdout.String(), Err: stderr.String()}
	switch {
	case errors.Is(ctx.Err(), context.DeadlineExceeded):
		res.Err += fmt.Sprintf("\n\ntimed out after %s — an infinite loop or a deadlock?", DefaultTimeout)
	case errors.Is(ctx.Err(), context.Canceled):
		res.Err += "\n\ncancelled."
	}
	return res, err
}

// cappedBuffer is an io.Writer that keeps at most maxOutput bytes and then
// records that it truncated, instead of growing without bound.
type cappedBuffer struct {
	buf       bytes.Buffer
	truncated bool
}

func (c *cappedBuffer) Write(p []byte) (int, error) {
	if room := maxOutput - c.buf.Len(); room > 0 {
		if len(p) > room {
			c.buf.Write(p[:room])
			c.truncated = true
		} else {
			c.buf.Write(p)
		}
	} else {
		c.truncated = true
	}
	// Report the full length: a short write would make exec treat the capped
	// stream as a broken pipe and fail the run for the wrong reason.
	return len(p), nil
}

func (c *cappedBuffer) String() string {
	if c.truncated {
		return fmt.Sprintf("%s\n… output truncated at %d KiB.", c.buf.String(), maxOutput/1024)
	}
	return c.buf.String()
}

func BuildArgs(e Exercise) []string {
	args := []string{}
	if e.Mode == "compile" {
		args = append(args, "run")
	} else {
		args = append(args, "test", "-v", "-race")
	}

	args = append(args, fmt.Sprintf("./%s", e.Path))
	return args
}
