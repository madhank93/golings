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

- A **type switch** inspects the type word inside the interface value and binds
  `x` to the concrete value, correctly typed, in each branch.

**Under the hood**

- An interface value carries (type, data). The switch compares that type word
  against each case; `default` runs when none matches. `any` is an alias for
  `interface{}` — the empty interface, which every type satisfies, and the
  reason `fmt.Println` takes anything.

**Common mistake**

- Using a plain assertion where the type is not guaranteed: `n := v.(int)`
  **panics** on a mismatch. The comma-ok form reports it instead:

```go
if n, ok := v.(int); ok { … }
```

**Key detail:** needing a type switch in your own API is often a sign a generic
(`generics1`) or a small interface would model it better — `any` throws away the
type information the compiler could have checked. Where it genuinely belongs:
`errors.As`, JSON decoding into `any`, and formatting.

**See also:** interfaces1 · switch1 (the expression switch it mirrors) ·
errors3 (`errors.As`) · generics1 · the [chapter](../README.md)

**References**

- Go spec — Type switches: https://go.dev/ref/spec#Type_switches ·
  Type assertions: https://go.dev/ref/spec#Type_assertions
- A Tour of Go — Type switches: https://go.dev/tour/methods/16
