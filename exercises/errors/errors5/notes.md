## errors5 — combine failures with errors.Join

```go
var errs []error
if len(p) < 8 {
    errs = append(errs, ErrTooShort)
}
if !strings.ContainsAny(p, "0123456789") {
    errs = append(errs, ErrNoDigit)
}
return errors.Join(errs...)
```

**Why it works**

- `errors.Join` (Go 1.20) bundles several errors into one value, and `errors.Is`
  matches **any** of them. So the caller can ask about each rule separately while
  the function reports them all at once.

**Under the hood**

- `Join` returns `nil` when the slice is empty or every element is `nil` — which
  is why the "collect, then join" shape needs no length check and returns a
  clean `nil` for a valid password. The result implements `Unwrap() []error`
  (note the slice), the multi-error form `errors.Is` and `errors.As` know how to
  traverse.

**Common mistake**

- Returning `errors.Join(errs...)` and then testing it with `==`. A joined error
  is a new value; only `errors.Is` reaches the members. Its `Error()` string is
  the members' messages separated by newlines.

**Key detail:** this is the right shape for validation, where reporting one
failure per round trip is a bad experience. It is also how a cleanup path
reports both the work error and a failed `Close`:
`defer func(){ err = errors.Join(err, f.Close()) }()`.

**See also:** errors2 (`%w` and `Is`) · errors3 (`As`) · defer2 (cleanup
errors) · the [chapter](../README.md)

**References**

- pkg.go.dev — errors.Join: https://pkg.go.dev/errors#Join
- Go 1.20 release notes: https://go.dev/doc/go1.20#errors
