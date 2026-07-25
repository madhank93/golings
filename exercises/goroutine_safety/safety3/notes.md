## safety3 — channel ownership

```go
ch := make(chan int, len(nums))
for _, n := range nums { ch <- n }
close(ch) // the OWNER (producer) closes, exactly once

for range 3 {
    wg.Go(func() {
        for v := range ch { total.Add(int64(v)) } // consumers only receive
    })
}
```

**Why it works**

- One goroutine owns the channel: it sends all values and `close`s it **once**.
  The three workers only **receive** (`range ch`), which ends cleanly for all of
  them when the channel closes.

**Nuance:** the ownership rule — **the sender closes, never a receiver**, and only
once. A consumer closing (or a second close) panics with "close of closed
channel". A single close broadcasts "no more values" to every ranging consumer
simultaneously, which is how you fan work out and shut it down.

**References**

- The Go Blog — Share Memory By Communicating: https://go.dev/blog/codelab-share
