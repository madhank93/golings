## logingest6 — slog middleware, and reading back a status you can't read

```go
func NewLogger(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level}))
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func WithLogging(next http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration", time.Since(start),
		)
	})
}
```

**Why it works**

- `slog` separates the *logger* from the *handler*. The handler decides format
  (JSON here) and destination and minimum level; the logger is just the API you
  call. Passing an `io.Writer` in rather than reaching for stderr is what makes
  the whole thing testable with a `bytes.Buffer`.
- Middleware is a function from `http.Handler` to `http.Handler`. Because
  `Server` from stage three is already an `http.Handler`, it wraps without
  modification — the service does not know it is being logged.
- Logging *after* `next.ServeHTTP` returns is what makes one record per request
  possible, and it is the only point at which the status and the duration are
  both known.

**Key detail:** `http.ResponseWriter` is write-only. There is no `Status()`
method, so middleware that wants the status has to intercept it — embed the real
`ResponseWriter` in a struct and shadow `WriteHeader`. Embedding matters: the
recorder inherits `Write` and `Header` for free and stays a valid
`ResponseWriter`, so you only override the one method you care about.

**Key detail:** the recorder starts at `http.StatusOK`, not at zero. A handler
may write a body and never call `WriteHeader` at all — `net/http` then sends 200
implicitly, and a recorder initialised to 0 logs `"status": 0`, a status code
that does not exist. Every dashboard built on that field is then quietly wrong
for the most common response in the service.

**Key detail:** `WriteHeader` must forward as well as record. An override that
only stores the code compiles, logs perfectly, and never sends the header — the
client gets a 200 no matter what the handler asked for. It is a bug that is
invisible in the logs, which is why the test asserts on `rec.Code` (what the
client saw) as well as on the log record.

**References**

- pkg.go.dev — log/slog: https://pkg.go.dev/log/slog
- Go blog — Structured Logging with slog: https://go.dev/blog/slog
- pkg.go.dev — http.ResponseWriter: https://pkg.go.dev/net/http#ResponseWriter
