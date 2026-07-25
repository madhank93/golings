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

- `ctx.Done()` is a channel that **closes** when `cancel()` is called. Selecting
  on it lets long-running work notice cancellation and stop promptly.

**Nuance:** cancellation is **cooperative** — the context can't kill your
goroutine; your code must *watch* `Done()` and return. `context.Background()` is
the empty root context you derive others from. Always call `cancel` (usually
`defer cancel()`) to release resources, even after normal completion.

**References**

- The Go Blog — Go Concurrency Patterns: Context: https://go.dev/blog/context
- Go by Example — Context: https://gobyexample.com/context
