## slog1 — handler levels

```go
h := slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
slog.New(h).Info(msg)
```

**Why it works**

- `log/slog` (Go 1.21) is structured logging in the stdlib. A handler has a
  **minimum level** and drops records below it. Pinned to `LevelWarn`, `Info`
  messages vanish; setting `Level: slog.LevelInfo` lets them through.

**Nuance:** levels are ordered `Debug < Info < Warn < Error`; the handler emits a
record only if its level is **≥** the threshold. `slog` writes **structured**
key/value records (`msg=... level=...`), not just text lines — machine-parseable
by default, which is the whole point over the old `log` package.

**References**

- pkg.go.dev — log/slog: https://pkg.go.dev/log/slog
- The Go Blog — Structured Logging with slog: https://go.dev/blog/slog
