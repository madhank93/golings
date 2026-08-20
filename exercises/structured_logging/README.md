# Structured logging (log/slog)

`log/slog` (Go 1.21) replaced a decade of "which logging library?" with an
answer in the standard library. A log record carries a message, a level, a time,
and typed **key/value attributes** — so a log line is queryable data rather than
a sentence you later regret writing a regex for.

The split that matters: a **`Logger`** is the API you call, a **`Handler`**
decides what happens to each record. Format, level filtering, destination and
enrichment all live in the handler, which is why swapping JSON for text — or
production for an in-memory test capture — changes one line.

## 1. Logger, handler, level

```go
h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
logger := slog.New(h)

logger.Info("server started", "port", 8080, "tls", true)
// time=… level=INFO msg="server started" port=8080 tls=true
```

Two handlers ship with the package: `NewTextHandler` (logfmt-ish, for humans)
and `NewJSONHandler` (for log pipelines). The four levels are `Debug`, `Info`,
`Warn`, `Error`, numbered −4, 0, 4, 8 — integers, so custom levels fit between
them.

**A handler drops every record below its level.** `slog1` is that: a handler
pinned to `LevelWarn` silently discards `Info`, and the test sees empty output.
Nothing errors; the record just never happens.

For a level you can change at run time, use `slog.LevelVar`:

```go
var lvl slog.LevelVar          // starts at Info
h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: &lvl})
lvl.Set(slog.LevelDebug)       // flip debug on in a live process
```

## 2. Attributes, and `With`

Attributes come in two spellings:

```go
logger.Info("done", "count", 3, "took", d)          // alternating k, v
logger.Info("done", slog.Int("count", 3), slog.Duration("took", d))  // typed
```

The alternating form is convenient and unchecked — an odd number of arguments
produces a `!BADKEY` entry instead of a compile error. The typed helpers
(`slog.String`, `Int`, `Bool`, `Duration`, `Any`, `Group`) are the safer choice
in code that matters.

`With` attaches attributes to **every** record from a logger:

```go
logger = logger.With("service", "api", "version", build)   // note the assignment
```

`slog2` is the trap: `With` **returns a new logger** and does not modify the
receiver. Calling `logger.With(...)` and throwing the result away compiles
cleanly and logs nothing extra. Same shape as `strings.Replace` or `append` —
in Go, a method that returns a value is telling you something.

Request-scoped logging is this plus `context`: put the derived logger in the
context in middleware, pull it out in handlers (`logingest7`).

## 3. Writing a handler

A `slog.Handler` is four methods:

```go
type Handler interface {
    Enabled(context.Context, Level) bool
    Handle(context.Context, Record) error
    WithAttrs([]Attr) Handler
    WithGroup(string) Handler
}
```

`slog3` implements the interesting one — `Handle` — over a slice, which is the
standard trick for asserting on logs in tests without parsing text:

```go
func (h MemHandler) Handle(_ context.Context, r slog.Record) error {
    *h.msgs = append(*h.msgs, r.Message)
    return nil
}
```

`Enabled` is called *before* the record is built, which is the mechanism behind
cheap disabled logging: when it returns false, the arguments are never
evaluated into attributes. `WithAttrs` and `WithGroup` must return a handler
carrying the extra state — returning `h` unchanged (as the exercise's stub does)
is fine for a test capture and wrong for a real handler.

A record's attributes are walked with `r.Attrs(func(a slog.Attr) bool { … })`.

## 4. Practical rules

- **Log at the boundary, not everywhere.** One line per request or per job,
  with the attributes that identify it — not one per function.
- **Never log and return the same error.** Pick one; the caller has context you
  do not (`errors` chapter).
- **Keys should be stable and low-cardinality**: `user_id`, `route`, `status`.
  Free-form message text is for humans; attributes are for machines.
- **Do not log secrets.** Structured output makes it easy to log a whole struct
  — including the token in it. Implement `LogValuer` on types that must redact:

```go
func (t Token) LogValue() slog.Value { return slog.StringValue("REDACTED") }
```

- **`slog.SetDefault(logger)`** routes the package-level `slog.Info` and the old
  `log` package through your handler — set it once in `main`.

## Gotchas

- **Records below the handler's level vanish silently** (`slog1`).
- **`With` returns a new logger** — assign it (`slog2`).
- **An odd number of alternating args** yields `!BADKEY`, not an error.
- **`slog.Any` on a large struct** serialises the whole thing, secrets included.
- **A custom handler must be safe for concurrent use** — several goroutines log
  through one handler.
- **`Enabled` is consulted first**, so guarding expensive attribute construction
  with `if logger.Enabled(ctx, slog.LevelDebug)` is worth it in hot paths.

## The exercises

- **slog1** — lower the handler's level so `Info` records survive.
- **slog2** — keep the logger `With` returns.
- **slog3** — implement `Handle` to capture records in memory.

## Source references

- [Go blog: Structured Logging with slog](https://go.dev/blog/slog) — the design,
  by its author
- [pkg.go.dev: log/slog](https://pkg.go.dev/log/slog) ·
  [slog.Handler](https://pkg.go.dev/log/slog#Handler) ·
  [LogValuer](https://pkg.go.dev/log/slog#LogValuer)
- [A Guide to Writing slog Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)

**Next: [reflection](../reflection/) →** — how `slog.Any` and `encoding/json`
inspect a value they were never told about.
