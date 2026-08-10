## logingest5 — graceful shutdown, and the context you must not derive

```go
serveErr := make(chan error, 1)
go func() { serveErr <- srv.Serve(ln) }()

select {
case err := <-serveErr:
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
case <-ctx.Done():
	shutCtx, cancel := context.WithTimeout(context.Background(), grace)
	defer cancel()
	return srv.Shutdown(shutCtx)
}
```

**Why it works**

- `srv.Serve` blocks until the server stops, so it cannot be the thing that
  watches `ctx`. Moving it to a goroutine turns "blocking call" into "value on a
  channel", which `select` can wait on alongside cancellation.
- `srv.Shutdown` stops the listener first, so no *new* connection is accepted,
  then waits for the handlers already running to return. That ordering is the
  whole of "graceful".
- `http.ErrServerClosed` is not a failure. `Serve` always returns a non-nil
  error, and this particular one means "somebody called Shutdown" — mapping it
  to `nil` is what lets the caller treat a non-nil return as a real problem.

**Key detail:** the shutdown context derives from `context.Background()`, never
from `ctx`. By the time that branch runs, `ctx` is *already cancelled* — a
context derived from it is born expired, `Shutdown` sees an exceeded deadline
immediately, and it closes active connections instead of waiting for them. The
result is a "graceful" shutdown that is exactly as abrupt as no shutdown at all,
and it passes every test that does not have a request in flight. That is why the
in-flight test exists.

**Key detail:** `make(chan error, 1)` — the buffer is not an optimisation. If
`ctx` is cancelled first, the `select` takes the other branch and nothing ever
receives from `serveErr`. On an unbuffered channel the `Serve` goroutine would
block on that send forever: a leaked goroutine per call, holding the whole
`http.Server` alive with it. One slot of buffer lets it send and exit.

**Key detail:** in the tests, every request goes through a client with a
`Timeout`. The listener is opened *before* `Serve` runs, so a `Serve` that never
accepts still completes the TCP handshake — the connection is established and
then silent. A bare `http.Get` against that blocks until the test binary's own
timeout, turning a clear failure into a ten-minute hang. Timeouts on clients in
tests are not belt-and-braces; they are what makes a broken implementation
*report* instead of *stall*.

**References**

- pkg.go.dev — http.Server.Shutdown: https://pkg.go.dev/net/http#Server.Shutdown
- pkg.go.dev — http.ErrServerClosed: https://pkg.go.dev/net/http#ErrServerClosed
- Go blog — Contexts and structs: https://go.dev/blog/context-and-structs
