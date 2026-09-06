## slog1 — the handler's level filters records

```go
h := slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
slog.New(h).Info(msg)
```

**Why it works**

- A handler drops every record below its minimum level. Pinned to `LevelWarn`,
  an `Info` call produces no output at all — no error, nothing written. Lowering
  it to `LevelInfo` lets the record through.

**Under the hood**

- The logger asks `handler.Enabled(ctx, level)` **before** building the record,
  so a filtered-out call skips attribute construction entirely. That is what
  makes leaving `Debug` calls in production code cheap.

**Common mistake**

- Expecting a level change to apply everywhere. The level belongs to the
  *handler*, so a logger built from a different handler keeps its own. For a
  level you can flip at run time, use a `*slog.LevelVar` — it is read on every
  check.

**Key detail:** the four levels are `Debug`, `Info`, `Warn`, `Error` at −4, 0, 4,
8 — plain integers, so custom levels fit between them. `slog.SetDefault(logger)`
routes the package-level `slog.Info` and the old `log` package through your
handler.

**See also:** slog2 (`With`) · slog3 (custom handlers) · logingest6 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — slog.HandlerOptions: https://pkg.go.dev/log/slog#HandlerOptions
- Go blog — Structured Logging with slog: https://go.dev/blog/slog
