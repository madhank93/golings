## interfaces3 — recover the concrete type

```go
func describe(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("int: %d", x)
    case string:
        return fmt.Sprintf("string: %s", x)
    default:
        return "unknown"
    }
}
```

**Why it works**

- A **type switch** (`v.(type)`) inspects the concrete type inside an interface
  value and binds `x` to it, correctly typed, in each case.

**Nuance:** for a single type use the **comma-ok assertion** `n, ok := v.(int)` —
`ok` is `false` instead of panicking when the type doesn't match. Plain
`v.(int)` (no `ok`) **panics** on a mismatch. `any` is just an alias for
`interface{}` — the empty interface every type satisfies.

**References**

- A Tour of Go — Type switches: https://go.dev/tour/methods/16
- Go by Example — Interfaces: https://gobyexample.com/interfaces
