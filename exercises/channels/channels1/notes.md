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

**Nuance:** you must `close` the channel or the `range` loop blocks forever
waiting for more. Only the **sender** should close, never the receiver. Sending
on a closed channel panics; receiving from one drains the buffer, then yields the
zero value with `ok == false`.

**References**

- A Tour of Go — Buffered Channels: https://go.dev/tour/concurrency/3
- Go by Example — Channel Buffering: https://gobyexample.com/channel-buffering
