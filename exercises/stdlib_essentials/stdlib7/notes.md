## stdlib7 — `omitzero` vs `omitempty`

One tag option changes:

```go
type Event struct {
	Name string    `json:"name"`
	At   time.Time `json:"at,omitzero"`
}
```

**Why it works**

- `omitempty` has a fixed, pre-generics definition of "empty": `false`, `0`, a
  nil pointer, a nil interface, and any array, slice, map, or string of length
  zero. A **struct** is not on that list, so a zero `time.Time` was encoded in
  full: `"at":"0001-01-01T00:00:00Z"`.
- `omitzero` (Go 1.24) asks the **type** instead. If it has an `IsZero() bool`
  method that answer wins; otherwise the field is compared against its type's
  zero value. `time.Time` has `IsZero`, so an unset timestamp disappears and a
  real one still marshals.

**Key detail:** the two options mean different things for scalars, and the
difference bites in the other direction too. `omitempty` on an `int` drops a
legitimate `0`, and on a `bool` drops a legitimate `false` — so a field that is
genuinely set to zero silently vanishes from the payload. The standard library
docs now recommend migrating `omitempty` to `omitzero` on bools, ints, uints,
floats, pointers, and interfaces for exactly that reason. Specifying both omits
the field when the value is empty **or** zero.

**References**

- pkg.go.dev — encoding/json Marshal: https://pkg.go.dev/encoding/json#Marshal
- Go 1.24 release notes: https://go.dev/doc/go1.24
- Go by Example — JSON: https://gobyexample.com/json
