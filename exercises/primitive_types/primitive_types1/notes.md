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

**Common mistake**

- Writing `storeIsOpen` on a line by itself, as the broken version does. Go has
  no bare-expression statement: a value must be used, assigned, or called. The
  compiler says `storeIsOpen (variable of type bool) is not used`.

**Key detail:** `bool` has exactly two values and **no implicit conversion**. `0`,
`""`, and `nil` are not falsy, and `1` is not true — an `if` condition must be a
genuine `bool`. That is why the idiom is `if err != nil`, never `if err`.

**See also:** if1 (conditions) · variables3 (zero value `false`) ·
the [chapter](../README.md)

**References**

- Go spec — Boolean types: https://go.dev/ref/spec#Boolean_types
- Go by Example — Values: https://gobyexample.com/values
