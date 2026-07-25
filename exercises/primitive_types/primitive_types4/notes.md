## primitive_types4 — a byte holds a character's code

```go
var b2 byte = 'a' // 'a' is the rune literal for code point 97
fmt.Println("representation for b2:", b2) // prints 97
```

**Why it works**

- `''` (empty) is an illegal rune literal. A single-quoted character like `'a'`
  is a **rune constant** whose numeric value (97) fits in a `byte`.

**Nuance:** `byte` is an alias for `uint8` — it stores a *number*, so printing it
shows `97`, not `a`. Single quotes make a character/rune; double quotes make a
`string`. They are different types.

**References**

- Go spec — Numeric types: https://go.dev/ref/spec#Numeric_types
- Go by Example — Values: https://gobyexample.com/values
