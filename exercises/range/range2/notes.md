## range2 — iterate a map's keys and values

```go
for name, phone := range phoneBook {
    fmt.Printf("%s has the %s phone\n", name, phone)
}
```

**Why it works**

- `range` over a map yields **key, value** each step (unlike a slice's
  index, value). So `name` is the key and `phone` is the value.

**Nuance:** map iteration order is **randomized** by design — never rely on it.
If you need a stable order, collect the keys into a slice and `sort` them, then
range over that.

**References**

- Go by Example — Range: https://gobyexample.com/range
