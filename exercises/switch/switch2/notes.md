## switch2 — a conditionless switch takes boolean cases

```go
switch {
case 0 > 1:
    ...
case 1 > 0:
    ...
}
```

**Why it works**

- `switch` with no expression is shorthand for a chain of `if/else`: each `case`
  must be a **boolean expression**. The broken `case:` was empty — every case
  needs a condition (or use `default`).

**Key detail:** `switch { ... }` (no value) is the idiomatic replacement for a long
`if/else if` ladder — often cleaner because each condition lines up as a `case`.

**References**

- Go by Example — Switch: https://gobyexample.com/switch
