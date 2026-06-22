# Sync primitives

The `sync` and `sync/atomic` packages coordinate goroutines that share memory.
Always test concurrent code with `go test -race` (this repo's `mise run test`
does) — it flags data races the eye misses.

- **sync1** — `sync.Mutex` to guard a shared counter
- **sync2** — `sync.Once` for run-exactly-once initialization
- **sync3** — `sync/atomic` for a lock-free counter

## Resources

- [Go by Example: Mutexes](https://gobyexample.com/mutexes)
- [Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)
- [pkg.go.dev: sync](https://pkg.go.dev/sync) · [sync/atomic](https://pkg.go.dev/sync/atomic)
