## reflect2 — walking a struct's fields

```go
func walk(x any, fn func(string)) {
    v := reflect.ValueOf(x)
    for i := 0; i < v.NumField(); i++ {
        if f := v.Field(i); f.Kind() == reflect.String {
            fn(f.String())
        }
    }
}
```

**Why it works**

- `reflect.ValueOf` gives a `reflect.Value` over the concrete value inside the
  interface; `NumField`/`Field(i)` enumerate a struct's fields, and the `Kind`
  check picks out the strings.

**Under the hood**

- Field *metadata* lives on the type side: `v.Type().Field(i)` returns a
  `StructField` with `Name`, `Tag` (`f.Tag.Get("json")` — this is exactly how
  `encoding/json` reads tags), and whether it is exported. **Unexported fields
  cannot be read via `Interface()` or set at all**, which is why JSON ignores
  lowercase fields.

**Common mistake**

- Calling `NumField` on a non-struct. It panics — real code checks
  `v.Kind() == reflect.Struct` first, and dereferences a pointer with `v.Elem()`
  before looking for fields.

**Key detail:** to *modify* through reflection the value must be addressable:
pass a pointer, call `.Elem()`, and check `CanSet()`. A `reflect.ValueOf(x)` over
a plain value holds a copy and can never be settable — Go's pass-by-value rule
showing through the reflection API.

**See also:** reflect1 (`Kind`) · stdlib1 (tags in practice) · structs1 ·
unsafe1 (layout without reflection) · the [chapter](../README.md)

**References**

- pkg.go.dev — reflect.Value: https://pkg.go.dev/reflect#Value ·
  StructTag: https://pkg.go.dev/reflect#StructTag
- Go blog — The Laws of Reflection: https://go.dev/blog/laws-of-reflection
