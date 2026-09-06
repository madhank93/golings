## process3 — "it failed" and "it never ran" are different

```go
err := exec.Command(name, args...).Run()
if err == nil {
    return 0, nil
}
var ee *exec.ExitError
if errors.As(err, &ee) {
    return ee.ExitCode(), nil   // ran, exited non-zero
}
return -1, err                  // never started
```

**Why it works**

- `*exec.ExitError` is returned **exactly when the child ran and exited
  non-zero**. Any other error — not on `PATH`, not executable, a bad working
  directory — means no process was ever created. Collapsing the two makes a
  shell that cannot tell "no match" from "grep is not installed".

**Under the hood**

- Exit codes are eight bits, 0–255. A process killed by a signal has no exit
  code of its own; Unix shells report `128 + signal`, and Go's `ExitCode()`
  returns **-1** for that case. `ee.ProcessState` carries the detail.

**Common mistake**

- `err.(*exec.ExitError)` instead of `errors.As`. A direct type assertion fails
  the moment anything wraps the error with `%w`, and wrapping is the norm.

**Key detail:** `exec.LookPath` performs the `PATH` search on its own and
returns `exec.ErrNotFound` — useful when you want to fail before doing any work,
rather than at the moment of `Run`.

**See also:** process1 (running it) · errors3 (`errors.Is`/`As`) · the
[chapter](../README.md)

**References**

- pkg.go.dev — exec.ExitError: https://pkg.go.dev/os/exec#ExitError
- pkg.go.dev — os.ProcessState.ExitCode: https://pkg.go.dev/os#ProcessState.ExitCode
