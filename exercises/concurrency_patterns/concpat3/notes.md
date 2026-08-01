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

- Each stage takes an input channel, launches a goroutine that transforms values
  onto a new output channel, and returns it. Chaining stages
  (`double(generate(...))`) wires them together; closing propagates from stage to
  stage.

**Key detail:** the golden rule of pipelines — **each stage closes its own output**
when its input is exhausted, so `close` ripples downstream and every `range`
terminates. Stages run concurrently and stream values as they arrive, rather than
materializing the whole result at each step.

**References**

- The Go Blog — Pipelines and cancellation: https://go.dev/blog/pipelines
