## slog3 — a custom slog.Handler

```go
func (h MemHandler) Handle(_ context.Context, r slog.Record) error {
    *h.msgs = append(*h.msgs, r.Message)
    return nil
}
```

**Why it works**

- A `slog.Handler` decides what happens to each record. Implementing `Handle`
  (the other three methods are stubs here) lets `MemHandler` capture every
  message into a slice — handy for asserting on logs in tests.

**Nuance:** the handler is the **backend** of slog; swap it to change output
without touching call sites (`TextHandler`, `JSONHandler`, or your own). `Handle`
receives a fully-built `slog.Record` — read `r.Message`, `r.Level`, `r.Time`, and
range `r.Attrs`. Storing via `*h.msgs` (a pointer) lets the value-receiver handler
mutate shared state.

**References**

- pkg.go.dev — slog.Handler: https://pkg.go.dev/log/slog#Handler
