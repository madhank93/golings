## logingest1 — a sentinel error the caller can match on

```go
func (e Event) Validate() error {
	switch {
	case e.Source == "":
		return fmt.Errorf("%w: source is empty", ErrInvalidEvent)
	case !slices.Contains(Levels, e.Level):
		return fmt.Errorf("%w: level %q is not one of %v", ErrInvalidEvent, e.Level, Levels)
	case e.Message == "":
		return fmt.Errorf("%w: message is empty", ErrInvalidEvent)
	case e.At.IsZero():
		return fmt.Errorf("%w: at is the zero time", ErrInvalidEvent)
	}
	return nil
}
```

**Why it works**

- `%w` wraps rather than formats. `fmt.Errorf("%w: …", ErrInvalidEvent)` builds a
  new error whose text names the field, while `errors.Is(err, ErrInvalidEvent)`
  still returns true — one error value carrying both halves of the contract.
- A single sentinel keeps the caller's job small. The stage-three HTTP handler
  will map *any* `ErrInvalidEvent` to `400 Bad Request` with one `errors.Is`
  check, instead of enumerating five separate error types.
- `slices.Contains` over a package-level `Levels` slice keeps the closed set in
  one place. Stage six's slog handler reads the same slice.

**Key detail:** `e.At.IsZero()`, not `e.At == time.Time{}`. Comparing
`time.Time` with `==` also compares the monotonic-clock reading and the
location pointer, so two times that represent the same instant can compare
unequal. `IsZero` asks the only question that is actually well-defined.

**Key detail:** this exercise is the first that spans two files in one package.
`go test ./exercises/capstone/logingest1/event_test.go` would fail with
`undefined: Event` — a file list compiles only the files named in it. Golings
marks the exercise `pkg = true` in `info.toml` and hands the toolchain the
directory instead.

**References**

- Go blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
- pkg.go.dev — errors.Is: https://pkg.go.dev/errors#Is
- pkg.go.dev — time.Time.IsZero: https://pkg.go.dev/time#Time.IsZero
