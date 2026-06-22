# Context

`context.Context` carries cancellation signals, deadlines, and request-scoped
values across API boundaries and goroutines. Convention: pass it as the first
argument, named `ctx`.

- **context1** — cancellation via `context.WithCancel` + `ctx.Done()`
- **context2** — timeouts via `context.WithTimeout`; return `ctx.Err()`
- **context3** — request-scoped values via `WithValue` / `ctx.Value`

## Resources

- [Go blog: Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [pkg.go.dev: context](https://pkg.go.dev/context)
- [Go by Example: Context](https://gobyexample.com/context)
