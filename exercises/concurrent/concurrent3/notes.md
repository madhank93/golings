## concurrent3 — send and receive over a channel

```go
go func() {
    messages <- "Hello"
    messages <- "World"
    close(messages)
}()
for msg := range messages { ... } // receives until closed
```

**Why it works**

- One goroutine **sends** two strings then closes the channel; the main goroutine
  **ranges** over it, receiving each value until the close ends the loop.

**Key detail:** an **unbuffered** channel (`make(chan string)`) is a synchronization
point — each send blocks until a receiver is ready, so the two goroutines hand
values off in lock-step. The sender must `close`, or the `range` would block
forever waiting for a third value.

**References**

- Go by Example — Channels: https://gobyexample.com/channels
- A Tour of Go — Channels: https://go.dev/tour/concurrency/2
