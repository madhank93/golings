## select1 — wait on multiple channels

```go
select {
case v := <-a:
    return v
case v := <-b:
    return v
}
```

**Why it works**

- `select` blocks until **one** of its cases can proceed, then runs that case. So
  `firstReady` returns whichever of `a` or `b` has a value first.

**Under the hood**

- The compiler lowers this to `runtime.selectgo`, which makes two passes: it
  polls the cases in a **randomly permuted** order looking for one that can
  proceed, and if none can, it enqueues the goroutine on *every* channel's wait
  queue and parks. The first channel to become ready wakes it and the runtime
  dequeues it from the others. A waiting `select` burns no CPU.

**Common mistake**

- Treating case order as priority — "check the urgent channel first". Ready cases
  are chosen uniformly at random. For real priority, try the urgent channel in a
  `select` with `default`, then fall back to the full `select`.

**Key detail:** every case's channel expression is evaluated once, up front,
before the wait — which matters as soon as one of them is `time.After(d)`. A
`select` with no ready case and no `default` blocks; an empty `select {}` blocks
forever; a case on a `nil` channel can never be chosen, which is how you switch
a case off.

**See also:** select2 (timeout) · select3 (`default`) · context1 (`ctx.Done()` as
a case) · the [select chapter](../README.md)

**References**

- Go spec — Select statements: https://go.dev/ref/spec#Select_statements
- src/runtime/select.go — `selectgo`:
  https://github.com/golang/go/blob/master/src/runtime/select.go
- Go by Example — Select: https://gobyexample.com/select
