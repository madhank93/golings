// process4
// Catch a termination signal instead of dying on it.

// I AM NOT DONE
package main_test

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

// ShutdownContext returns a context that is cancelled when SIGINT or SIGTERM
// arrives, or when parent is cancelled, whichever comes first. The returned
// function releases the handler and cancels the context; callers defer it.
//
// The default disposition of SIGTERM is to kill the process where it stands:
// no deferred functions, no flush, no goodbye. Registering a handler replaces
// that for the signals you name — and only those, so SIGKILL still cannot be
// caught, which is why "kill -9 leaves a lock file behind" is a real bug class.
//
// Releasing the handler matters as much as installing it. A process that keeps
// swallowing SIGTERM after it has stopped listening is a process nobody can
// stop politely.
func ShutdownContext(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	// FIXME: three things are wrong here.
	//
	//  1. SIGTERM is the signal a supervisor, an init system or `kill` sends;
	//     SIGINT alone catches Ctrl-C and nothing else.
	//  2. Nothing ever receives from ch, so the signal is registered and then
	//     ignored. Cancel ctx when one arrives.
	//  3. The returned function must release the handler as well as cancel,
	//     or a later SIGTERM keeps being swallowed by a channel nobody reads.
	//
	// signal.NotifyContext does all three in one call — or build it yourself
	// with a goroutine, which is worth doing once to see what it is made of.
	return ctx, cancel
}

const helperEnv = "GOLINGS_PROCESS4_HELPER"

// The child installs the handler, says so, then waits to be told to stop. The
// "ready" line is printed only once the handler is in place — signalling
// before that point would kill it with the default disposition and this test
// would be flaky rather than wrong.
func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		t.Skip("not the child")
	}
	base, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ctx, stop := ShutdownContext(base)
	defer stop()

	fmt.Println("ready")
	<-ctx.Done()
	fmt.Println("caught")
	os.Exit(0)
}

func TestShutdownContextCatchesSIGTERM(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("signals are a Unix mechanism; Windows has no SIGTERM to send")
	}
	t.Setenv(helperEnv, "1")

	cmd := exec.Command(os.Args[0], "-test.run=^TestHelperProcess$")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill() //nolint:errcheck

	sc := bufio.NewScanner(stdout)
	if !sc.Scan() || sc.Text() != "ready" {
		t.Fatalf("the child never became ready (got %q)", sc.Text())
	}

	if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
		t.Fatal(err)
	}
	if !sc.Scan() {
		t.Fatal("the child died on SIGTERM instead of catching it")
	}
	if got := sc.Text(); got != "caught" {
		t.Errorf("child printed %q, want %q", got, "caught")
	}
	if err := cmd.Wait(); err != nil {
		t.Errorf("child exited with %v, want a clean exit", err)
	}
}

func TestShutdownContextFollowsItsParent(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	ctx, stop := ShutdownContext(parent)
	defer stop()

	select {
	case <-ctx.Done():
		t.Fatal("the context is done before anything happened")
	default:
	}

	cancelParent()
	select {
	case <-ctx.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("cancelling the parent did not cancel the shutdown context")
	}
	if !errors.Is(ctx.Err(), context.Canceled) {
		t.Errorf("ctx.Err() = %v, want context.Canceled", ctx.Err())
	}
}

func TestShutdownContextStopReleasesTheHandler(t *testing.T) {
	ctx, stop := ShutdownContext(context.Background())
	stop()

	select {
	case <-ctx.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("stop() must cancel the context as well as release the handler")
	}
}
