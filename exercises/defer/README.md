# Defer

`defer` schedules a call to run when the surrounding function returns. Deferred
calls run in LIFO order and execute on every return path — the idiomatic way to
guarantee cleanup (`defer f.Close()`).

- **defer1** — LIFO order (and shaping a named return value)
- **defer2** — guaranteed cleanup on every return path

## Resources

- [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12)
- [Go by Example: Defer](https://gobyexample.com/defer)
- [Effective Go: Defer](https://go.dev/doc/effective_go#defer)
