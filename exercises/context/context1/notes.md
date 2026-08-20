## context1 — cancellation with context

```go
func countUntilCancelled(ctx context.Context) int {
    for {
        select {
        case <-ctx.Done():
            return count // caller cancelled
        case <-time.After(time.Millisecond):
            count++
        }
    }
}
ctx, cancel := context.WithCancel(context.Background())
```

**Why it works**

- `WithCancel` returns a context and a `cancel` function. Calling `cancel()`
  closes `ctx.Done()`, so the `select` takes that case and the loop returns.

**Under the hood**

- `Done()` hands back a channel that is **closed**, never sent to — which is why
  every watcher is released at once and why later receives return immediately.
  Internally the `cancelCtx` also holds its children; `cancel()` closes its own
  channel, records `Err()`, then walks the children cancelling each. Cancellation
  flows down the tree only: cancelling a child never affects its parent.

**Common mistake**

- Assuming cancellation stops the goroutine. It is **cooperative** — nothing in
  the runtime interrupts anything. A loop that never selects on `ctx.Done()` (or
  never checks `ctx.Err()` between chunks of CPU work) runs to completion no
  matter how many times `cancel` is called.

**Key detail:** `defer cancel()` at every call site, even when you expect the work
to finish first. It releases the child from the parent's children set right away;
`go vet`'s `lostcancel` check exists because forgetting it is a silent leak. Once
cancelled, a context stays cancelled — there is no reset.

**See also:** select1 (the `select` this is built on) · context2 (a deadline
doing the cancelling) · concpat3 (the leak cancellation prevents) · the
[context chapter](../README.md)

**References**

- Go blog — Context: https://go.dev/blog/context
- src/context/context.go — `cancelCtx`, `propagateCancel`:
  https://github.com/golang/go/blob/master/src/context/context.go
- pkg.go.dev — context: https://pkg.go.dev/context
