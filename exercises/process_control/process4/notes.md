## process4 — catch the signal instead of dying on it

```go
ctx, stop := signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
defer stop()
<-ctx.Done()
```

Or, built by hand — worth doing once:

```go
ch := make(chan os.Signal, 1)          // buffered: delivery never blocks
signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
go func() {
    defer signal.Stop(ch)
    select {
    case <-ch:
    case <-parent.Done():
    }
    cancel()
}()
```

**Why it works**

- The default disposition of SIGTERM is to terminate the process immediately:
  no deferred functions, no flush, no goodbye. `signal.Notify` replaces that for
  the signals you name, turning a fatal event into a value on a channel — which
  is what makes graceful shutdown possible at all.

**Under the hood**

- The runtime's signal handler does the one thing that is safe in that context:
  a non-blocking send. If nothing is ready to receive, **the signal is dropped**
  — which is why the channel must be buffered, and why the docs are emphatic
  about it. One slot is enough: a second SIGTERM before you have handled the
  first says nothing new.

**Common mistake**

- Never calling `signal.Stop` (or the `stop` from `NotifyContext`). The handler
  stays installed, later signals keep landing in a channel nobody reads, and the
  process becomes one that only `kill -9` can stop.

**Key detail:** SIGKILL and SIGSTOP cannot be caught, blocked or ignored — the
kernel enforces it. Everything a process wants to do on the way out must happen
on SIGTERM, which is exactly why supervisors send SIGTERM first and SIGKILL only
after a grace period.

**See also:** logingest5 (graceful HTTP shutdown) · context1 (cancellation) ·
process3 (what a signalled child reports) · the [chapter](../README.md)

**References**

- pkg.go.dev — os/signal: https://pkg.go.dev/os/signal
- pkg.go.dev — signal.NotifyContext: https://pkg.go.dev/os/signal#NotifyContext
