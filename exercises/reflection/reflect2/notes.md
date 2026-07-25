## reflect2 — walk a struct's fields

```go
v := reflect.ValueOf(x)
for i := 0; i < v.NumField(); i++ {
    if v.Field(i).Kind() == reflect.String {
        fn(v.Field(i).String())
    }
}
```

**Why it works**

- `reflect.ValueOf(x)` gives a `reflect.Value` you can introspect. `NumField()` /
  `Field(i)` iterate a struct's fields; checking `Kind() == reflect.String` and
  reading `.String()` visits only the string fields (`Name`, `City` — not `Age`).

**Nuance:** this field-walking is exactly how `encoding/json`, validators, and ORMs
work generically. Caveats: `NumField` panics if `x` isn't a struct, and you can
only read **exported** fields' values via reflection. It's powerful but slower and
unchecked — use it for framework-level code, not everyday logic.

**References**

- The Go Blog — The Laws of Reflection: https://go.dev/blog/laws-of-reflection
- pkg.go.dev — reflect: https://pkg.go.dev/reflect
