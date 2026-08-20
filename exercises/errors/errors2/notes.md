## errors2 — wrap with %w

```go
var ErrNotFound = errors.New("not found")

func lookup(id string) error {
    if id == "" {
        return fmt.Errorf("lookup failed: %w", ErrNotFound)
    }
    return nil
}
```

**Why it works**

- `%w` stores the original error **inside** the new one instead of flattening it
  to text, so `errors.Is(err, ErrNotFound)` still finds the sentinel behind the
  added context.

**Under the hood**

- `fmt.Errorf` with a `%w` returns a `*wrapError` carrying the message and an
  `Unwrap() error` method. `errors.Is` calls `Unwrap` repeatedly, comparing each
  link with `==` (or calling a custom `Is` method), until it matches or the chain
  ends.

**Common mistake**

- Testing with `err == ErrNotFound`. That works only while nothing has wrapped
  it — one caller adding context breaks every such comparison. `errors.Is` is
  the check that survives wrapping.

**Key detail:** wrapping is an **API decision**. `%w` makes the inner error part
of your package's contract, so callers may write `errors.Is` against it; `%v`
deliberately hides it. Choose per error, not per habit.

**See also:** errors6 (a chain broken by `%v`) · errors3 (`errors.As`) ·
errors5 (`Join`) · the [chapter](../README.md)

**References**

- Go blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
- pkg.go.dev — errors.Is: https://pkg.go.dev/errors#Is
