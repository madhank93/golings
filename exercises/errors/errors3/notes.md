## errors3 — custom error types + errors.As

```go
type ValidationError struct{ Field string }
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s is required", e.Field)
}
var ve *ValidationError
errors.As(err, &ve) // extracts the concrete type from the chain
```

**Why it works**

- Any type with an `Error() string` method **is** an error. `errors.As` searches
  the wrap chain for one matching `*ValidationError` and, when found, sets `ve` to
  it so you can read `ve.Field`.

**Key detail:** `errors.Is` answers *"is this a specific error value?"*; `errors.As`
answers *"is there an error of this type, and give it to me."* Use `As` when you
need the typed error's **fields**, not just its identity.

**References**

- errors.As: https://pkg.go.dev/errors#As
- The Go Blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
