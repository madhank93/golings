## tcp4 — one goroutine per connection

```go
for {
    conn, err := ln.Accept()
    if err != nil { /* … */ }

    wg.Add(1)
    go func() {
        defer wg.Done()
        defer conn.Close()
        handle(conn, c)
    }()
}
```

**Why it works**

- The accept loop stays single-threaded and the *work* moves off it. A slow
  client then holds one goroutine — a few kilobytes of stack — instead of the
  whole server. This is why Go took over network services: the blocking code you
  would write anyway becomes concurrent by adding one word.

**Under the hood**

- The runtime multiplexes those goroutines onto a handful of OS threads with
  epoll/kqueue underneath. A blocking `conn.Read` parks its goroutine and frees
  the thread, so ten thousand idle connections cost ten thousand small stacks,
  not ten thousand threads.

**Common mistake**

- Sharing state across handlers without a lock. `Counter` is touched by every
  connection at once, and an unguarded `c.n++` is a data race: it may even
  produce plausible-looking numbers. `go test -race` — which golings always runs
  — is what turns that into a failure you can see.

**Key detail:** unbounded goroutines are a denial-of-service vector; a real
server bounds them with a semaphore channel or a worker pool. The `WaitGroup`
here does something different and also necessary: it stops `Serve` from
returning while a handler is still writing.

**See also:** concurrency_patterns2 (fan-in) · safety1 (guarding shared state) ·
sync1 (`Mutex`) · the [chapter](../README.md)

**References**

- pkg.go.dev — sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
- The Go Blog — Go's work-stealing scheduler: https://go.dev/blog/go15gc
