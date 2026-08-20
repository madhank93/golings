## errors3 — custom error types and errors.As

```go
type ValidationError struct{ Field string }

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid field %q", e.Field)
}

var ve *ValidationError
if errors.As(err, &ve) {
    use(ve.Field) // the data survives the wrapping
}
```

**Why it works**

- Any type with an `Error() string` method is an `error`, so a struct can carry
  structured detail — here the offending field. `errors.As` walks the wrap chain
  looking for something assignable to the pointer you pass, and assigns it.

**Under the hood**

- The method is declared on `*ValidationError`, so the **pointer** is what
  satisfies `error` — value receivers would make both `T` and `*T` work, pointer
  receivers only `*T`. That is why the code returns `&ValidationError{…}` and
  matches into a `*ValidationError`.

**Common mistake**

- Passing the wrong thing to `errors.As`. It needs a **pointer to** the target
  variable: `errors.As(err, &ve)`, not `errors.As(err, ve)`. Passing a
  non-pointer panics.

**Key detail:** sentinel or type? A sentinel (`errors2`) answers *which* failure
happened; a custom type answers *with what data*. Use the type when the caller
needs a field — the retry-after duration, the invalid field, the HTTP status.

**See also:** errors2 (sentinels) · interfaces3 (type switches) ·
methods1 (method sets) · errors5 · the [chapter](../README.md)

**References**

- pkg.go.dev — errors.As: https://pkg.go.dev/errors#As
- Go blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
