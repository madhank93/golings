## stdlib7 — omitzero vs omitempty

```go
type Event struct {
    Name string    `json:"name"`
    At   time.Time `json:"at,omitzero"` // was omitempty
}
```

**Why it works**

- `omitempty` cannot drop a struct — it has no notion of an "empty" struct — so a
  zero `time.Time` marshals as `"0001-01-01T00:00:00Z"`. `omitzero` (Go 1.24)
  asks the type whether it is the zero value, and a zero time disappears.

**Under the hood**

- `omitempty` predates most of Go's type system: it drops `""`, `0`, `false`, and
  nil/empty maps, slices and pointers — a fixed list. `omitzero` compares against
  the type's zero value, using the type's own `IsZero() bool` method when it has
  one (`time.Time` does).

**Common mistake**

- Reaching for `*time.Time` to get the same effect. A pointer does work with
  `omitempty`, at the cost of a nil check at every use site — `omitzero` gets
  there without changing the field's type.

**Key detail:** they answer different questions. `omitempty` is right for "drop
the empty slice"; `omitzero` for "drop the unset value". An empty-but-non-nil
slice is *empty* but not *zero*, so the two can disagree — use the one that
matches your intent.

**See also:** stdlib1 (tags) · variables3 (zero values) · logingest1 (event
model with tags) · the [chapter](../README.md)

**References**

- Go 1.24 release notes — encoding/json: https://go.dev/doc/go1.24#encoding-json
- pkg.go.dev — encoding/json: https://pkg.go.dev/encoding/json
