# Concurrency patterns

Composable building blocks for concurrent Go, all built from goroutines +
channels. "Don't communicate by sharing memory; share memory by communicating."

- **concpat1** — worker pool: N workers consume one jobs channel
- **concpat2** — fan-in: merge many channels into one
- **concpat3** — pipeline: chain stages via channels

## Resources

- [Go blog: Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go blog: Concurrency Patterns](https://go.dev/blog/io2012-videos)
- [Go by Example: Worker Pools](https://gobyexample.com/worker-pools)
