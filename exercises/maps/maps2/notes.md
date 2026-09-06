## maps2 — build a map with a literal

```go
func ages() map[string]int {
    return map[string]int{
        "John": 30,
        "Ana":  21,
    }
}
```

**Why it works**

- A composite literal states the full type and its entries in one expression, so
  the map is allocated and filled together. `map{}` is missing the key and value
  types, which are part of the type's name.

**Common mistake**

- Dropping the trailing comma on the last entry. Go's formatter puts each entry
  on its own line, and every line — including the last — needs its comma before
  the closing brace, or the build fails.

**Key detail:** a literal is the right choice for a fixed set of entries;
`make(map[K]V, hint)` is for maps you fill at run time, where the hint
preallocates buckets and saves rehashing. Values can be any type, including
another map (`map[string]map[string]int`) or a struct.

**See also:** maps1 (`make`) · maps3 (reading back) · mapspkg1 (`maps.Keys`) ·
the [chapter](../README.md)

**References**

- Go spec — Composite literals: https://go.dev/ref/spec#Composite_literals
- Go by Example — Maps: https://gobyexample.com/maps
