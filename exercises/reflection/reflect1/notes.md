## reflect1 — inspect a value's kind

```go
func describe(x any) string {
    return reflect.TypeOf(x).Kind().String()
}
// describe(42) == "int"; describe([]int{1}) == "slice"
```

**Why it works**

- `reflect.TypeOf(x)` recovers the dynamic type inside an `any`; `.Kind()` reduces
  it to a category (`int`, `string`, `slice`, `map`, `struct`, ...). So `describe`
  reports what *kind* of thing it was handed at run time.

**Key detail:** **Type vs Kind** — `Type` is the specific type (`main.Celsius`); `Kind`
is the underlying category (`float64`). Reflection trades compile-time safety for
run-time flexibility, so reach for it only when you genuinely can't know types
ahead of time (serializers, ORMs) — a type switch is clearer when you can.

**References**

- The Go Blog — The Laws of Reflection: https://go.dev/blog/laws-of-reflection
- pkg.go.dev — reflect: https://pkg.go.dev/reflect
