# Reflection

Reflection is the ability of a program to inspect and manipulate values whose
types it does not know at compile time. In Go it lives in one package,
`reflect`, and it powers the things that *cannot* be written any other way:
`encoding/json`, database mappers, `fmt`'s `%v`, validators, and deep-equality
in tests.

It is also the part of Go where the compiler stops helping you. Every mistake
becomes a run-time panic, the code is slower, and the tooling cannot see what
you are doing. Most programs should reach for an interface or a type parameter
first — and then, occasionally, this.

## 1. The three laws

Rob Pike's framing, from *The Laws of Reflection*:

1. **Reflection goes from interface value to reflection object.**
   `reflect.TypeOf(x)` and `reflect.ValueOf(x)` take an `any` and give you a
   `Type` and a `Value`.
2. **Reflection goes from reflection object back to interface value.**
   `v.Interface()` returns an `any` you can type-assert.
3. **To modify a reflection object, the value must be addressable and
   settable** — which means you must have passed a *pointer* in.

```ascii
     any (type, data)
        │  reflect.ValueOf / TypeOf
        ▼
  reflect.Value / reflect.Type      <- inspect: Kind, NumField, Field(i)
        │  v.Interface()
        ▼
     any again  -> type assert
```

## 2. Type vs Kind

```go
reflect.TypeOf(42).Kind()          // reflect.Int    — the category
reflect.TypeOf(Celsius(1)).Name()  // "Celsius"      — the named type
reflect.TypeOf(Celsius(1)).Kind()  // reflect.Float64 — its underlying kind
```

`reflect1` returns `reflect.TypeOf(x).Kind().String()`, which gives `"int"`,
`"string"`, `"slice"`, `"map"`. The distinction is the one that catches people:
**Kind is the structural category** (there are 26 of them), **Type is the
specific named type**. A `type Celsius float64` has kind `Float64` and type
`Celsius`, and code that switches on kind will handle every defined type built
on `float64` at once.

## 3. Walking a struct

```go
v := reflect.ValueOf(x)
for i := 0; i < v.NumField(); i++ {
    f := v.Field(i)
    if f.Kind() == reflect.String {
        fn(f.String())
    }
}
```

`reflect2`. `NumField` and `Field(i)` only work when the kind is `Struct` —
calling them on anything else panics, so real code checks
`v.Kind() == reflect.Struct` first.

The field's *metadata* comes from the type side: `v.Type().Field(i)` gives a
`StructField` with the `Name`, the `Tag` (`f.Tag.Get("json")` — this is how
JSON tags are read), and whether it is exported.

Which brings the biggest restriction: **unexported fields can be read only as
opaque values and never set.** `f.Interface()` on an unexported field panics.
That is why `encoding/json` ignores lowercase fields — not a policy decision, a
consequence of reflection's rules.

## 4. Modifying: addressability

```go
func doubleIt(p any) {
    v := reflect.ValueOf(p).Elem()    // Elem() dereferences the pointer
    if v.CanSet() {
        v.SetInt(v.Int() * 2)
    }
}

n := 21
doubleIt(&n)     // n is 42
```

`reflect.ValueOf(n)` holds a **copy**, so it can never be settable — the same
pass-by-value rule as everywhere else in Go, showing through. Pass a pointer,
call `.Elem()`, and check `CanSet()` before writing. Forgetting is a panic:
`reflect: reflect.Value.SetInt using unaddressable value`.

## 5. When to use it — and what to use instead

Legitimate uses are narrow and structural: serialisation, ORMs, dependency
wiring, test helpers like `reflect.DeepEqual`, and framework glue that must work
for types it has never seen.

Everything else has a better answer now:

| Instead of reflection | Use |
|---|---|
| "any type with these methods" | an interface |
| "same algorithm, many types" | generics (Go 1.18+) |
| "which concrete type is this?" | a type switch (`interfaces3`) |
| comparing two values in a test | `slices.Equal`, `maps.Equal`, or `cmp.Diff` |

Reflection also costs: allocations on every `ValueOf`, no inlining, and no
compile-time checking. Fine at startup or per request; measurable in a loop.

## Gotchas

- **`NumField`/`Field` panic on non-structs.** Check `Kind()` first.
- **Unexported fields cannot be read via `Interface()` or set** — ever.
- **A non-pointer `Value` is never settable**; pass `&x` and call `.Elem()`.
- **`Kind` is not `Type`.** Named types share the kind of their underlying type.
- **`reflect.DeepEqual` is not a general equality check** — it compares
  unexported fields, treats `nil` and empty slices as different, and is slow.
  Prefer typed comparisons.
- **Reflection code fails at run time**, so it needs test coverage that
  compiled code would not.

## The exercises

- **reflect1** — report a value's `Kind` as a string.
- **reflect2** — walk a struct's fields and act on the string ones.

## Source references

- [Go blog: The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [pkg.go.dev: reflect](https://pkg.go.dev/reflect) ·
  [reflect.Kind](https://pkg.go.dev/reflect#Kind) ·
  [StructTag](https://pkg.go.dev/reflect#StructTag)
- [Go blog: JSON and Go](https://go.dev/blog/json) — reflection in production use

**Next: [unsafe_pkg](../unsafe_pkg/) →** — one step further out, where even the
type system's guarantees stop.
