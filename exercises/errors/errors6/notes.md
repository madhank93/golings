## errors6 — %w keeps the chain; %v breaks it

```go
level1 := fmt.Errorf("read file: %w", ErrNotFound)
level2 := fmt.Errorf("load config: %w", level1) // %w, not %v
level3 := fmt.Errorf("startup: %w", level2)
errors.Is(level3, ErrNotFound) // true — three layers deep
```

**Why it works**

- Each `%w` links a new message onto the existing error, building a chain.
  `errors.Is` walks the **whole** chain, so `ErrNotFound` is found three layers
  down. A single `%v` anywhere flattens that layer to plain text and severs the
  chain below it.

**Nuance:** `%w` vs `%v` is the switch between "wrap (still matchable)" and
"format (text only)". Use `%w` when a caller might need `errors.Is`/`As`; use
`%v` only when you deliberately want to hide the underlying error.

**References**

- The Go Blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
- errors.Is: https://pkg.go.dev/errors#Is
