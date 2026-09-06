## sync2 — run exactly once with sync.Once

```go
func (c *Config) Load() {
    c.once.Do(func() {
        c.loads++
    })
}
```

**Why it works**

- `sync.Once`'s `Do` runs its function **the first time only** — every later call
  (from any goroutine) returns without running it again. So across 50 concurrent
  `Load()` calls, `loads` ends at 1.

**Under the hood**

- `Once` is a `done` flag plus a mutex. `Do` begins with an **atomic load** of the
  flag, so after the first call it costs less than taking a lock. If the flag is
  unset, the slow path locks, re-checks, runs `f`, and sets the flag in a
  `defer`. Callers arriving mid-run block on that mutex until the first run
  finishes, which is why nobody can observe a half-initialised value — a plain
  `if !inited { init() }` cannot promise that.

**Common mistake**

- Using `Once` for work that can fail. The flag is set in a `defer`, so a panic
  or an error inside `f` **still marks it done** and no retry ever happens.
  `Once` means "attempted once". Store the error and decide, or manage the retry
  yourself.

**Key detail:** the zero value is ready to use, and a `Once` must never be copied
after first use. When the point is producing a value rather than running a side
effect, prefer the Go 1.21 helpers: `sync.OnceValue(f)` returns a function that
computes once and caches, with `OnceFunc` and `OnceValues` alongside.

**See also:** sync1 (the mutex underneath) · di1 (constructor injection instead
of lazy globals) · the [sync chapter](../README.md)

**References**

- src/sync/once.go: https://github.com/golang/go/blob/master/src/sync/once.go
- pkg.go.dev — sync.Once: https://pkg.go.dev/sync#Once ·
  OnceValue: https://pkg.go.dev/sync#OnceValue
