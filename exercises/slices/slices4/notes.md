## slices4 — index and sub-slice within bounds

```go
names[0]   // single element → "John"
names[0:2] // sub-slice     → ["John" "Maria"]
names[2:4] // sub-slice     → ["Carl" "Peter"]
```

**Why it works**

- A single index (`names[0]`) yields **one element**; a range (`names[0:2]`)
  yields a **slice**. Keeping every index inside `0..len` avoids a panic.

**Nuance:** two different result types from one operator — `s[i]` is the element
type (`string`), `s[i:j]` is a slice (`[]string`). And the high bound may equal
`len(s)` (that's the end); only going *past* `len` panics.

**References**

- Go by Example — Slices: https://gobyexample.com/slices
