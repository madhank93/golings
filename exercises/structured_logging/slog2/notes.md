## slog2 — logger.With returns a new logger

```go
logger = logger.With("service", "api") // reassign — With returns a NEW logger
logger.Info(msg)                        // now every line carries service=api
```

**Why it works**

- `logger.With(attrs...)` returns a **new** logger that stamps those attributes on
  every record. It does **not** mutate the receiver, so discarding its result
  (the broken code) means the attribute never appears. Assigning it back fixes it.

**Nuance:** this immutable, copy-returning style is deliberate — you build a
base logger, then derive request- or component-scoped loggers with `.With(...)`
and pass them down, without one derivation affecting another. Same pattern as
`context.WithValue`.

**References**

- pkg.go.dev — slog.Logger.With: https://pkg.go.dev/log/slog#Logger.With
