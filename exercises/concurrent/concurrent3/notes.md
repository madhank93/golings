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

**Under the hood**

- The channel is unbuffered, so there is nowhere to store the value: the runtime
  copies it straight from the sender's stack to the receiver's stack and wakes
  the receiver. Whoever arrives first parks in the channel's send or receive
  queue. That handoff is also a memory-model edge — everything the sender did
  before the send is visible to the receiver after it.

**Common mistake**

- Forgetting the `close`. `range` ends on *closed*, not on *empty*, so the loop
  blocks forever waiting for a third value and the program dies with
  `all goroutines are asleep - deadlock!`.

**Key detail:** this is the alternative to concurrent2 — instead of two goroutines
sharing a variable and guarding it, one hands the value over and the handoff *is*
the synchronization. No shared state, nothing to lock. Only the sender closes;
a receiver that closes breaks the producer.

**See also:** channels1 (buffering) · channels2 (directional types) ·
safety3 (who is allowed to close) · the [goroutines chapter](../README.md)

**References**

- Go spec — Channel types: https://go.dev/ref/spec#Channel_types
- The Go Memory Model — channels: https://go.dev/ref/mem#chan
- A Tour of Go — Channels: https://go.dev/tour/concurrency/2
