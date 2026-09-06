## logingest8 — pprof on a mux you own, and the finished service

```go
func WithPprof(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(PprofPrefix, pprof.Index)
	mux.HandleFunc(PprofPrefix+"cmdline", pprof.Cmdline)
	mux.HandleFunc(PprofPrefix+"profile", pprof.Profile)
	mux.HandleFunc(PprofPrefix+"symbol", pprof.Symbol)
	mux.HandleFunc(PprofPrefix+"trace", pprof.Trace)
	return mux
}

func NewService(store *Store, log *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", NewServer(store))
	WithPprof(mux)
	return WithLogging(mux, log)
}
```

**Why it works**

- `net/http/pprof` is written around a global. Its `init()` registers the
  handlers on `http.DefaultServeMux`, which is why the usual advice is the
  one-line blank import — and why that advice does nothing for a mux you built
  yourself.
- Registering the five entry points explicitly gives you the profiles without
  adopting whatever else has been attached to the default mux by a transitive
  dependency. `pprof.Index` covers the named profiles (`heap`, `goroutine`,
  `allocs`, `block`, `mutex`, `threadcreate`) on its own, so five registrations
  is the whole surface.
- `ServeMux` matches the **longest** pattern, not the first registered. That is
  what lets `mux.Handle("/", …)` sit under `"/debug/pprof/"` without shadowing
  it, and it is why the order of the two registrations does not matter.
- Wrapping outermost puts every request through the logger — profile fetches
  included, which is what you want when someone is pulling a 30-second CPU
  profile off a production box.

**Key detail:** the alternative you will see in the wild is
`mux.Handle("/debug/pprof/", http.DefaultServeMux)`, forwarding the whole
default mux under that prefix. It is shorter and it works, but it exposes
everything anything else has ever registered on the default mux, reachable
under `/debug/pprof/`. Explicit registration is a few more lines for a surface
you can actually enumerate.

**Key detail:** these endpoints are not safe to expose publicly. `/debug/pprof/
cmdline` leaks the process arguments, `profile` will happily consume 30 seconds
of CPU per request, and the goroutine dump describes your internals in detail.
In a real deployment they belong on a separate listener bound to localhost or to
an internal interface — the same handlers, a different `http.Server`.

**And that is the capstone.** An event model with a sentinel error, a
concurrency-safe store, method-pattern routing, a bounded worker pool, graceful
shutdown, structured logging, a fuzz-tested parser, and runtime profiling —
assembled by `NewService` into one `http.Handler` that stage five's `Serve` can
run until it is told to stop.

**References**

- pkg.go.dev — net/http/pprof: https://pkg.go.dev/net/http/pprof
- Go blog — Profiling Go Programs: https://go.dev/blog/pprof
- pkg.go.dev — http.ServeMux pattern precedence: https://pkg.go.dev/net/http#ServeMux
