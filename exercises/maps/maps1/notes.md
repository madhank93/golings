## maps1 — a map type names both halves

```go
m := make(map[string]int)
m["John"] = 30
m["Ana"] = 21
fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
```

**Why it works**

- `map` alone is not a type: it needs the key type and the value type,
  `map[string]int`. `make` then allocates the hash table, and `m[k] = v` inserts.

**Under the hood**

- A map value is a **pointer to a runtime hash table**, so passing a map to a
  function shares it — inserts and deletes made inside are visible to the
  caller, unlike appends to a slice. The zero value is `nil`: readable, but
  writing to it panics with `assignment to entry in nil map`.

**Common mistake**

- `var m map[string]int` followed by `m["k"] = 1`. Declaring gives you a `nil`
  map; `make` or a literal gives you a usable one.

**Key detail:** the key type must be **comparable** — strings, numbers, bools,
pointers, and structs or arrays of comparable fields. Slices, maps and functions
cannot be keys, because `==` is not defined for them.

**See also:** maps2 (literals) · maps3 (missing keys) · safety2 (concurrent
maps) · the [chapter](../README.md)

**References**

- Go spec — Map types: https://go.dev/ref/spec#Map_types
- Go blog — Go maps in action: https://go.dev/blog/maps
