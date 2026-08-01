## maps1 — declare a map's key and value types

```go
m := make(map[string]int)
m["John"] = 30
```

**Why it works**

- `make(map)` is incomplete — a map type is `map[KeyType]ValueType`. Here keys
  are `string` (names) and values are `int` (ages), so `map[string]int`.

**Key detail:** a map must be **initialized** with `make` (or a literal) before you
write to it. The zero value of a map is `nil`; reading a `nil` map is fine
(returns the zero value), but **writing** to one **panics**.

**References**

- Go by Example — Maps: https://gobyexample.com/maps
