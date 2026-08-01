## maps2 — build a map with a literal

```go
func ages() map[string]int {
    return map[string]int{"John": 30, "Ana": 21}
}
```

**Why it works**

- `map{}` is not valid syntax. A **map literal** names the full type then lists
  `key: value` pairs: `map[string]int{"John": 30, "Ana": 21}` — declared and
  filled in one expression.

**Key detail:** the literal both allocates and initializes, so (unlike `make`) you
get a ready-to-use, non-nil map. Reading a **missing** key never errors — it
returns the value type's zero value (`0` for `int`).

**References**

- Go by Example — Maps: https://gobyexample.com/maps
