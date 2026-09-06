## tcp3 — the accept loop, and how it ends

```go
for {
    conn, err := ln.Accept()
    if err != nil {
        if errors.Is(err, net.ErrClosed) {
            return nil // ln was closed: this is the shutdown
        }
        return err
    }
    handle(conn)
}
```

**Why it works**

- A server is a loop around `Accept`. Closing the listener is how you stop it:
  the blocked `Accept` returns immediately with `net.ErrClosed`, which is a
  shutdown signal wearing an error's clothes. Any other error is real.

**Under the hood**

- `http.Server.Serve` is this loop with two extras: it retries temporary errors
  with a backoff, and it hands each connection to a goroutine. Everything else
  in it is HTTP.

**Common mistake**

- Returning the error unconditionally, so every clean shutdown reports a
  failure — or, worse, `continue`-ing on all errors, which turns a closed
  listener into a spin loop that burns a core.

**Key detail:** `errors.Is(err, net.ErrClosed)` is the check to use.
`err.Error() == "use of closed network connection"` compares against a string
the standard library is free to reword, and did.

**See also:** tcp4 (a goroutine per connection) · errors3 (`errors.Is`) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — net.ErrClosed: https://pkg.go.dev/net#pkg-variables
- Go source — http.Server.Serve: https://cs.opensource.google/go/go/+/master:src/net/http/server.go
