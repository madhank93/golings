## arrays2 — one element type per array

```go
names := [4]string{"John", "Maria", "Carl", "Peter"}
```

**Why it works**

- The broken literal mixed a number (`10`) into a `[4]string`. Every element of
  an array shares the **one** declared element type, so the stray `int` had to
  become a string.

**Key detail:** Go has no heterogeneous arrays. If you genuinely need mixed types,
that's a struct (named fields) or `[]any` (an explicit escape hatch you rarely
want) — not an array.

**References**

- Go by Example — Arrays: https://gobyexample.com/arrays
