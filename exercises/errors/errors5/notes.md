## errors5 — bundle errors with errors.Join

```go
var errs []error
if len(p) < 8            { errs = append(errs, ErrTooShort) }
if !strings.ContainsAny(p, "0123456789") { errs = append(errs, ErrNoDigit) }
return errors.Join(errs...) // one error, or nil if the slice is empty
```

**Why it works**

- `errors.Join` (Go 1.20+) combines several errors into one. `errors.Is` then
  matches **any** of the joined errors, so both `ErrTooShort` and `ErrNoDigit` are
  reachable.

**Nuance:** `errors.Join` returns `nil` when every argument is `nil` (or the slice
is empty) — so a fully valid password naturally yields `nil`. It's the idiomatic
way to report **all** validation failures at once instead of stopping at the
first.

**References**

- errors.Join: https://pkg.go.dev/errors#Join
- Go by Example — Errors: https://gobyexample.com/errors
