## primitive_types1 — a bool is just a value you can reassign

```go
storeIsOpen := true
// ...first check runs...
storeIsOpen = false
// ...now the !storeIsOpen check runs...
```

**Why it works**

- The store starts open, so the first `if storeIsOpen` fires. To make the second
  check (`if !storeIsOpen`) fire too, you reassign `storeIsOpen = false` between
  them.

**Nuance:** `bool` has exactly two values, `true` and `false`, and no implicit
conversion — Go won't treat `0`/`1` or a non-empty string as a bool the way some
languages do. A condition must be a genuine `bool`.

**References**

- Go by Example — Values: https://gobyexample.com/values
