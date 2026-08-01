## modern4 — WaitGroup.Go

```go
for _, n := range nums {
    wg.Go(func() {
        total.Add(int64(n))
    })
}
wg.Wait()
```

**Why it works**

- Go 1.25's `wg.Go(f)` launches `f` in its own goroutine and handles the `Add(1)`
  and `Done()` for you — replacing the error-prone `wg.Add(1); go func(){ defer
  wg.Done(); ... }()` dance. The broken code's bare `go` never called `Add`, so
  `Wait` returned before the work ran.

**Key detail:** this removes the two classic `WaitGroup` bugs — forgetting `Add`
(Wait returns too early) and forgetting `Done` (Wait blocks forever). Note the
result uses `atomic.Int64` because many goroutines add concurrently.

**References**

- pkg.go.dev — sync.WaitGroup.Go: https://pkg.go.dev/sync#WaitGroup.Go
