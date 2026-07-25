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

**Nuance:** `Once` is the concurrency-safe way to do lazy, one-time init
(singletons, config loading, connection setup) without a mutex-and-flag dance.
`Do` also **blocks** concurrent callers until the first run finishes, so nobody
sees a half-initialized value.

**References**

- sync.Once: https://pkg.go.dev/sync#Once
