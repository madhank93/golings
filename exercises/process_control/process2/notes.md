## process2 — write to a child's stdin

```go
cmd := exec.Command(name, args...)
cmd.Stdin = strings.NewReader(input)
out, err := cmd.Output()
```

**Why it works**

- Setting `Stdin` to any `io.Reader` makes `os/exec` do the plumbing: it creates
  the pipe, copies the reader into it in a goroutine, and — the part that
  matters — **closes the pipe when the reader is exhausted**. That close is the
  EOF the child is waiting for.

**Under the hood**

- This is one half of a shell pipeline. `a | b` is two processes with `b`'s
  stdin wired to `a`'s stdout; in Go that is `b.Stdin, _ = a.StdoutPipe()`, and
  byox's shell course is largely this idea plus job control.

**Common mistake**

- Using `StdinPipe` and never closing it. The child reads until EOF, the EOF
  never comes, and both processes wait for each other — a deadlock that looks
  like a hang, not a failure. With `cmd.Stdin = reader` the close is automatic.

**Key detail:** a child that writes more than a pipe buffer (~64 KiB on Linux)
blocks until someone drains it. Feed a large stdin *and* read stdout only after
`Wait`, and you deadlock; `Output` avoids it by draining concurrently.

**See also:** process1 (capturing stdout) · files2 (`bufio.Scanner`) · di1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — exec.Cmd.Stdin: https://pkg.go.dev/os/exec#Cmd
- pkg.go.dev — exec.Cmd.StdinPipe: https://pkg.go.dev/os/exec#Cmd.StdinPipe
