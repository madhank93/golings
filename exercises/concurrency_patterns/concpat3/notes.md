## concpat3 — pipeline

```go
func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out) // close downstream when upstream is done
    }()
    return out
}
// double(generate(1, 2, 3)) → 2, 4, 6
```

**Why it works**

- Each stage owns exactly one channel — the one it created — and closes it when
  its input runs dry. So the shutdown cascades: the generator closes, `double`'s
  `range` ends, `double` closes, and the consumer's loop ends. No coordination
  code anywhere.

**Under the hood**

- The channels are unbuffered, so each stage runs only as fast as the next stage
  consumes: a send parks until the downstream receive arrives. That is
  backpressure for free — the generator cannot race ahead and build an unbounded
  queue in memory.

**Common mistake**

- Leaving the pipeline early. `for v := range double(...) { if v > 4 { break } }`
  leaves the stage goroutine parked on a send nobody will ever receive — the
  classic Go goroutine leak: silent, permanent, invisible to tests. Give every
  stage a `ctx` and select on `ctx.Done()` alongside the send.

**Key detail:** `defer close(out)` as the first line inside the goroutine is
sturdier than closing at the end — it survives an early `return` from a stage
that hits an error. Stages compose by nesting (`filter(double(generate(…)))`)
precisely because each takes and returns `<-chan T`.

**See also:** context1 (the cancellation that fixes the leak) · concpat2
(fan-in) · channels2 (direction) · safety3 (ownership) · the
[patterns chapter](../README.md)

**References**

- Go blog — Pipelines and cancellation: https://go.dev/blog/pipelines
- Go blog — Go Concurrency Patterns: https://go.dev/talks/2012/concurrency.slide
- pkg.go.dev — context: https://pkg.go.dev/context
