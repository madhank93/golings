## channels1 — buffered channels

```go
ch := make(chan int, len(vals)) // buffer of len(vals)
for _, v := range vals { ch <- v }
close(ch)
for v := range ch { out = append(out, v) } // drains until closed
```

**Why it works**

- `make(chan int, n)` gives the channel a **buffer** of `n`, so `n` sends succeed
  without a receiver waiting. After `close(ch)`, `for v := range ch` receives every
  buffered value and then stops.

**Under the hood**

- The buffer is a ring: an array of `n` slots plus a send index, a receive index,
  a count, and a mutex. A send with room copies into the slot at `sendx` and
  returns; a send with no room parks the goroutine in the channel's send queue.
  With capacity 0 there is no array at all, which is why the unbuffered version
  of this single-goroutine loop deadlocks — there is no second goroutine to hand
  the value to.

**Common mistake**

- Sizing the buffer to "make it fast". Capacity is a backpressure decision, not a
  speed dial: a large buffer only delays the moment a slow consumer becomes
  visible, at the cost of memory and latency.

**Key detail:** you must `close` or the `range` blocks forever. Only the sender
closes. Sending on a closed channel panics, closing twice panics, and receiving
from a closed channel drains the buffer first, then yields the zero value with
`ok == false` forever.

**See also:** concurrent3 (unbuffered handoff) · channels2 (direction) ·
select3 (non-blocking send/receive) · the [channels chapter](../README.md)

**References**

- Go spec — Channel types: https://go.dev/ref/spec#Channel_types
- src/runtime/chan.go — `hchan`, `chansend`, `chanrecv`:
  https://github.com/golang/go/blob/master/src/runtime/chan.go
- A Tour of Go — Buffered Channels: https://go.dev/tour/concurrency/3
