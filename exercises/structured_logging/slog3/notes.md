## slog3 — implementing a Handler

```go
func (h MemHandler) Handle(_ context.Context, r slog.Record) error {
    *h.msgs = append(*h.msgs, r.Message)
    return nil
}
```

**Why it works**

- A `slog.Handler` decides what happens to each record. `Handle` receives the
  `slog.Record` — message, level, time, attributes — and this one appends the
  message to a slice, so a test can assert on logs without parsing text.

**Under the hood**

- The interface is four methods: `Enabled` (called first, so a disabled level
  costs nothing), `Handle`, and `WithAttrs`/`WithGroup`, which must return a
  handler carrying the added state. Returning `h` unchanged is fine for a test
  capture and wrong for a real handler.

**Common mistake**

- Holding the slice by value instead of by pointer. `MemHandler` is used as a
  value, so a `msgs []string` field would append into a copy — the same trap as a
  spy with a value receiver (`mock1`). The `*[]string` is what makes the recording
  visible.

**Key detail:** a handler must be **safe for concurrent use** — several goroutines
log through one handler. This test double is not, which is acceptable in a
single-goroutine test and would need a mutex anywhere else. Attributes are walked
with `r.Attrs(func(a slog.Attr) bool { ... })`.

**See also:** slog1 · slog2 · mock1 (the same capture idea) · safety1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — slog.Handler: https://pkg.go.dev/log/slog#Handler
- A Guide to Writing slog Handlers: https://github.com/golang/example/blob/master/slog-handler-guide/README.md
