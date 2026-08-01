## channels2 — channel directions

```go
func produce(out chan<- int, n int) { ... close(out) } // send-only
func consume(in <-chan int) []int   { for v := range in ... } // receive-only
```

**Why it works**

- `chan<- int` is a **send-only** channel; `<-chan int` is **receive-only**. A
  plain `chan int` converts to either. Declaring the direction in the signature
  documents intent and lets the compiler enforce it.

**Key detail:** direction is a **compile-time** guard — `produce` literally cannot
receive, and `consume` cannot send or close. This makes producer/consumer roles
unmistakable and catches misuse (like a consumer closing the channel) at build
time.

**References**

- Go by Example — Channel Directions: https://gobyexample.com/channel-directions
