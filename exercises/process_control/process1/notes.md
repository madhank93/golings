## process1 — run a child, capture its output

```go
out, err := exec.Command(name, args...).Output()
if err != nil {
    return "", err
}
return strings.TrimRight(string(out), "\n"), nil
```

**Why it works**

- `exec.Command` builds a `*Cmd` — a description of a process, not a running
  one. `Output` starts it, collects stdout, waits for it to exit, and hands back
  the bytes. Nothing runs until you call `Run`, `Start`, `Output` or
  `CombinedOutput`.

**Under the hood**

- The tests need a real child process, so they re-run **the test binary itself**
  with an environment marker (`os.Args[0]` plus `-test.run=^TestHelperProcess$`).
  That is the pattern `os/exec`'s own tests use: no shell, no coreutils, and it
  works identically on every platform `go test` runs on.

**Common mistake**

- Reaching for `CombinedOutput` and wondering why stderr is in the result.
  `Output` captures stdout alone — and it populates `ExitError.Stderr` for you,
  which is the diagnostic you actually want when the command fails.

**Key detail:** `Start` without a matching `Wait` leaves a zombie: the child has
exited but its status is never reaped. `Run`, `Output` and `CombinedOutput` all
wait for you.

**See also:** process2 (feeding it stdin) · process3 (its exit code) · cli1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — os/exec: https://pkg.go.dev/os/exec
- Go source — exec_test.go's helper-process pattern: https://cs.opensource.google/go/go/+/master:src/os/exec/exec_test.go
