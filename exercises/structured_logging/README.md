# Structured Logging (log/slog)

`log/slog` (Go 1.21) is the standard library's structured, leveled logger and
the backbone of modern Go services. Records carry typed key/value attributes,
handlers control format and filtering, and you can implement your own handler to
route logs anywhere.

- **slog1** — handler levels: why `Info` disappears under a `Warn` filter
- **slog2** — `logger.With(...)` returns a new logger you must keep
- **slog3** — implementing a custom `slog.Handler`

## Resources

- [log/slog package](https://pkg.go.dev/log/slog)
- [Go blog: Structured Logging with slog](https://go.dev/blog/slog)
