# Concurrency (goroutines)

A goroutine is a lightweight thread started with `go f()`. Goroutines
communicate over **channels** (`chan T`) rather than shared memory, and
`sync.WaitGroup` waits for a group to finish. This topic introduces the basics;
`channels`, `select`, `sync`, `context`, and `concurrency_patterns` go deeper.

## Resources

- [A Tour of Go: Goroutines](https://go.dev/tour/concurrency/1)
- [Go by Example: Goroutines](https://gobyexample.com/goroutines) · [Channels](https://gobyexample.com/channels)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
