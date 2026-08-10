# Capstone — a concurrent log-ingest service

The capstone is one program built across eight stages. Each stage is its own
runnable package that ships with the previous stages already solved, so you can
start anywhere and reset a stage without losing the rest.

By the end you have a service that accepts log events over HTTP, validates
them, batches them through a worker pool, stores them safely under concurrent
access, logs with `log/slog`, shuts down cleanly on cancellation, and exposes
`net/http/pprof`.

These are the first **package-mode** exercises: each one spans several files in
a directory, and golings compiles the whole directory rather than a single
file.

- **logingest1** — the `Event` model: struct tags, a sentinel error, `%w`
- **logingest2** — a concurrency-safe store: `sync.RWMutex`, generics
- **logingest3** — the HTTP surface: `ServeMux` routing, JSON decode/encode
- **logingest4** — a worker pool: goroutines, channels, `sync.WaitGroup`
- **logingest5** — cancellation and graceful shutdown: `context`, `http.Server.Shutdown`
- **logingest6** — structured logging: `log/slog`, a handler, request middleware
- **logingest7** — testing the service: `httptest`, subtests, fuzzing, `b.Loop`
- **logingest8** — profiling in production: `net/http/pprof` on a custom mux

## Resources

- [Go blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [pkg.go.dev: net/http](https://pkg.go.dev/net/http)
- [pkg.go.dev: log/slog](https://pkg.go.dev/log/slog)
- [Go blog: Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
