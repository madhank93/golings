# Process control

A shell is a program that runs other programs, and `os/exec` plus `os/signal`
is all the machinery that takes. The Go part is small; the ideas are Unix ones —
a child inherits your environment and file descriptors, reports one byte of exit
status, and can be interrupted by a signal it may or may not be allowed to
catch.

Three questions run through the chapter: what did the child *print*, what did it
*read*, and how did it *end*.

## 1. Running a child

```go
out, err := exec.Command(name, args...).Output()
```

`process1`. `exec.Command` builds a description of a process; nothing runs until
`Run`, `Start`, `Output` or `CombinedOutput`. `Output` starts it, collects
stdout, waits, and returns the bytes — and populates `ExitError.Stderr` when it
fails, which is the diagnostic you want.

`Start` gives you a running child and hands the waiting back to you. Every
`Start` needs a matching `Wait`, or the child becomes a zombie: exited, but its
status never reaped.

**Where the tests get a child.** They re-run the test binary itself —
`os.Args[0]` with `-test.run=^TestHelperProcess$` and an environment marker —
which is the pattern `os/exec`'s own test suite uses. No shell, no coreutils,
identical behaviour on every platform.

## 2. Standard input, and the pipeline

```go
cmd.Stdin = strings.NewReader(input)
```

`process2`. Give `Stdin` any `io.Reader` and `os/exec` creates the pipe, copies
into it, and **closes it when the reader is exhausted**. That close is the EOF
the child is waiting for; a `StdinPipe` you never close is a deadlock, not an
error.

```
   Filter():  strings.Reader ──► [pipe] ──► child stdin
                                            child stdout ──► [pipe] ──► Output()
```

Wiring two children together is the same idea and is what a shell's `|` does:

```go
b.Stdin, _ = a.StdoutPipe()
```

Pipe buffers are finite (~64 KiB on Linux). A child that writes more than that
blocks until someone drains it, so "write all the input, then read all the
output" deadlocks on a large enough payload. `Output` avoids it by draining
concurrently.

## 3. Exit status

```go
var ee *exec.ExitError
if errors.As(err, &ee) {
    return ee.ExitCode(), nil   // it ran, and it failed — that is a result
}
return -1, err                  // it never started — that is an error
```

`process3`. These are different failures and code that conflates them is a shell
that cannot tell "no match" from "grep is not installed". `*exec.ExitError` is
returned exactly when a process was created and exited non-zero.

Exit codes are eight bits, 0–255. A process killed by a signal has no exit code
of its own: shells report `128 + signal` (so 143 for SIGTERM), and Go's
`ExitCode()` returns **-1**, with the detail in `ProcessState`.

Use `errors.As`, not a type assertion — wrapped errors are the norm.

## 4. Signals

```go
ctx, stop := signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
defer stop()
<-ctx.Done()
```

`process4`. The default disposition of SIGTERM is to terminate the process where
it stands — no defers, no flush. `signal.Notify` (which `NotifyContext` wraps)
replaces that for named signals, turning a fatal event into a value on a
channel.

```
   kill -TERM ──► runtime handler ──► non-blocking send ──► your channel
                                            │
                                            └─ nobody ready? the signal is DROPPED
```

Hence the two rules: **the channel must be buffered** (one slot is enough), and
**`signal.Stop`** must eventually run, or the process keeps swallowing signals
it no longer handles.

SIGKILL and SIGSTOP cannot be caught, blocked or ignored — the kernel enforces
it. That is why supervisors send SIGTERM first and SIGKILL only after a grace
period, and why anything that must happen on the way out happens on SIGTERM.

Signals are a Unix mechanism; Windows has no SIGTERM to send, so `process4`'s
child test skips there and the rest of the chapter runs everywhere.

## Gotchas

- **Every `Start` needs a `Wait`.** Otherwise: zombie.
- **`Output` is stdout only**; `CombinedOutput` merges stderr in.
- **An unclosed stdin pipe is a hang**, not an error.
- **Don't read stdout only after `Wait`** on a chatty child — the pipe fills and
  both processes stop.
- **`errors.As(&ee)`**, never `err.(*exec.ExitError)`.
- **`ExitCode()` is -1 for a signalled child**, not `128+n`.
- **Buffer the signal channel**, or the signal is dropped.
- **Call `signal.Stop`** when you are done listening.
- **Never build a command with `sh -c "… " + userInput`** — that is shell
  injection. `exec.Command` takes an argv slice precisely so no shell parses it.

## The exercises

- **process1** — run a child and capture what it printed.
- **process2** — feed a child on stdin and read its answer.
- **process3** — tell a failed command apart from one that never ran.
- **process4** — catch SIGTERM instead of dying on it.

## Where this goes next

byox's shell course is 76 stages, of which 19 are tagged *process control* and
14 *terminal i/o*: run a command, wire up `|`, redirect with `>` and `2>`, keep
a job in the background, forward Ctrl-C to the foreground job. Every one of them
is this chapter plus a parser — and the terminal half (raw mode via
`golang.org/x/term`) is the piece golings deliberately leaves to it.

## Source references

- [pkg.go.dev: os/exec](https://pkg.go.dev/os/exec) ·
  [exec.ExitError](https://pkg.go.dev/os/exec#ExitError) ·
  [os/signal](https://pkg.go.dev/os/signal)
- [Go source: exec_test.go](https://cs.opensource.google/go/go/+/master:src/os/exec/exec_test.go)
  — the helper-process pattern these exercises use
- [signal(7)](https://man7.org/linux/man-pages/man7/signal.7.html) — dispositions,
  and what cannot be caught

**End of the on-ramp.** Three chapters, twelve exercises, and the three stdlib
areas a language roadmap has no reason to cover and a systems project cannot
start without. Next stop is a project, not another exercise:
[byox](https://github.com/madhank93/build-your-own-x).
