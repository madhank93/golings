## maps3 — read, insert, and delete map keys

```go
phone := phoneBook["Ana"]      // read by key
phoneBook["Laura"] = "+11 ..." // insert / update
delete(phoneBook, "John")      // remove a key
```

**Why it works**

- Indexing with the right key returns its value; assigning to a key inserts or
  overwrites; `delete` removes a pair and shrinks `len`.

**Nuance:** reading a missing key returns the zero value, which is
indistinguishable from a key whose value *is* the zero value. When you must know
whether a key exists, use the **comma-ok** form: `v, ok := m[k]` — `ok` is
`false` when the key is absent.

**References**

- Go by Example — Maps: https://gobyexample.com/maps
