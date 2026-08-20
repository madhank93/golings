## channels2 — channel directions

```go
func produce(out chan<- int, n int) { ... close(out) } // send-only
func consume(in <-chan int) []int   { for v := range in ... } // receive-only
```

**Why it works**

- `chan<- int` is a **send-only** channel; `<-chan int` is **receive-only**. A
  plain `chan int` converts to either. Declaring the direction in the signature
  documents intent and lets the compiler enforce it.

**Under the hood**

- The restriction is purely a type-system view of the same `hchan`; there is no
  runtime cost and no second channel. The conversion is one-way — you cannot
  recover a bidirectional `chan T` from a `<-chan T`, which is what makes the
  guarantee hold all the way down a call chain.

**Common mistake**

- Returning a bidirectional `chan T` from a constructor-like function. The caller
  can then close a channel it does not own, and a producer still sending panics
  with `send on closed channel`. Return `<-chan T`; the fan-in and pipeline
  exercises do exactly that.

**Key detail:** direction turns the "only the owner closes" convention into a
**compile error** — `consume` literally cannot close or send. Prefer directional
parameters everywhere by default; the signature then says which end of the pipe
each side is holding.

**See also:** channels1 (buffering) · safety3 (the panic this prevents) ·
concpat2 / concpat3 (`<-chan T` return types) · the
[channels chapter](../README.md)

**References**

- Go spec — Channel types: https://go.dev/ref/spec#Channel_types
- Effective Go — Channels: https://go.dev/doc/effective_go#channels
- Go by Example — Channel Directions: https://gobyexample.com/channel-directions
