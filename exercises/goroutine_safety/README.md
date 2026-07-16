# Goroutine Safety

Sharing data across goroutines demands discipline. These exercises drill the
tools and rules that keep concurrent code correct under the race detector:
read/write locks for read-heavy state, `sync.Map` for concurrent maps, and the
channel-ownership rule. Run them with `go test -race`.

- **safety1** — `sync.RWMutex`: concurrent readers, exclusive writers
- **safety2** — `sync.Map` for safe concurrent map access
- **safety3** — channel ownership: only the owner closes, exactly once

## Resources

- [Go blog: Share Memory By Communicating](https://go.dev/blog/codelab-share)
- [sync.RWMutex](https://pkg.go.dev/sync#RWMutex) · [sync.Map](https://pkg.go.dev/sync#Map)
- [Data Race Detector](https://go.dev/doc/articles/race_detector)
