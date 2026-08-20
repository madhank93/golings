## slog2 — With returns a new logger

```go
logger := slog.New(slog.NewTextHandler(w, nil))
logger = logger.With("service", "api") // keep the result
logger.Info(msg)                        // service=api appears
```

**Why it works**

- `With` does not modify the receiver — it returns a **new** logger whose handler
  carries the extra attributes. Discarding the return value discards the
  attributes with it.

**Under the hood**

- Internally `With` calls `handler.WithAttrs(...)`, which returns a handler
  holding the pre-formatted attributes. That is why `With` is cheap to call once
  at startup and expensive to call per log line: the formatting happens once, at
  `With` time.

**Common mistake**

- The bare `logger.With("service", "api")` statement — it compiles, it allocates,
  and it changes nothing. Same family as `strings.Replace` or `append` without an
  assignment: a method returning a value is telling you something.

**Key detail:** attributes can be given as alternating key/value pairs or as typed
helpers (`slog.String`, `slog.Int`). The alternating form is unchecked — an odd
count produces a `!BADKEY` entry rather than a compile error.

**See also:** slog1 (levels) · slog3 (handlers) · context3 (request-scoped
loggers) · logingest6 · the [chapter](../README.md)

**References**

- pkg.go.dev — slog.Logger.With: https://pkg.go.dev/log/slog#Logger.With
- Go blog — Structured Logging with slog: https://go.dev/blog/slog
