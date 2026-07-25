## errors2 — wrap with %w, match with errors.Is

```go
var ErrNotFound = errors.New("not found")

return fmt.Errorf("lookup %q: %w", id, ErrNotFound) // %w wraps
errors.Is(err, ErrNotFound)                          // true, through the wrap
```

**Why it works**

- `%w` in `fmt.Errorf` **wraps** the sentinel while adding context. `errors.Is`
  then walks the wrap chain and finds `ErrNotFound` even under the extra message.

**Nuance:** use `%w` (not `%v`) when you want callers to still match the
underlying error. A **sentinel** — a package-level `var Err... = errors.New(...)`
— is the value callers compare against with `errors.Is`, never by string.

**References**

- The Go Blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
- errors.Is: https://pkg.go.dev/errors#Is
