## reflect1 — a value's Kind

```go
func describe(x any) string {
    return reflect.TypeOf(x).Kind().String()
}
```

**Why it works**

- `reflect.TypeOf` turns an interface value into a `reflect.Type`, and `Kind`
  reports its structural category — `int`, `string`, `slice`, `map` — as one of
  26 constants.

**Under the hood**

- **Kind is not Type.** A `type Celsius float64` has kind `Float64` and type
  `Celsius`; switching on kind therefore handles every defined type built on the
  same underlying type at once. `Type.Name()` gives the declared name,
  `Type.String()` the package-qualified one.

**Common mistake**

- Calling `reflect.TypeOf(nil)`, which returns a **nil** `Type` — so `.Kind()`
  panics. Any code taking `any` from outside needs that guard.

**Key detail:** reflection is the compiler's last resort and yours too. "Which
concrete type is this?" is usually a type switch (`interfaces3`); "the same
algorithm over many types" is generics. Reflection is for code that must work
with types it has never seen — encoders, ORMs, `fmt`.

**See also:** reflect2 (walking fields) · interfaces3 (type switches) ·
generics1 · stdlib1 (reflection in `encoding/json`) · the [chapter](../README.md)

**References**

- Go blog — The Laws of Reflection: https://go.dev/blog/laws-of-reflection
- pkg.go.dev — reflect.Kind: https://pkg.go.dev/reflect#Kind
