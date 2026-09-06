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

- Closing announces "no more sends". Only the sender can honestly promise that,
  so the producer closes once and the three workers simply range until the
  channel is drained and closed.

**Under the hood**

- `close` flips a flag on the channel and wakes every goroutine parked in its
  wait queues. Closing an already-closed channel panics with
  `close of closed channel`, and a send after close panics with
  `send on closed channel` — both are unrecoverable in practice because they fire
  in whichever worker lost the race, not in the code that made the mistake.

**Common mistake**

- Letting a consumer close "when it is done". With three workers, the first close
  succeeds and the second kills the program. It is also load-dependent: with one
  worker the bug never shows.

**Key detail:** make the rule structural instead of remembered — hand consumers a
receive-only `<-chan T` and `close` will not compile. When several producers
share one output, nobody owns it individually, so hoist ownership to a closer
goroutine: `go func(){ wg.Wait(); close(out) }()` (concpat2). To stop producers
from the consumer side, cancel a context — never close the data channel.

**See also:** channels2 (directional types) · concpat2 (the closer goroutine) ·
concurrent3 (sender closes) · the [safety chapter](../README.md)

**References**

- Go spec — Close: https://go.dev/ref/spec#Close
- Go blog — Pipelines and cancellation: https://go.dev/blog/pipelines
- pkg.go.dev — sync.WaitGroup.Go: https://pkg.go.dev/sync#WaitGroup.Go
