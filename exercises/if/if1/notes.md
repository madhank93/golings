## if1 — return the larger value

```go
func bigger(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Why it works**

- The `if` handles one case and returns; whatever reaches the last line is the
  other case. No `else` is needed, and none should be written.

**Common mistake**

- Reaching for a ternary. Go has none — `a > b ? a : b` does not compile. Since
  Go 1.21 there is a built-in for exactly this: `return max(a, b)`.

**Key detail:** every path must return. Wrapping the second `return` in an `else`
compiles too, but the flat version is the house style: handle the special case,
leave, and keep the main path unindented down the left margin.

**See also:** if2 (chained conditions) · switch3 (when there are more than two
branches) · the [chapter](../README.md)

**References**

- Effective Go — If: https://go.dev/doc/effective_go#if
- pkg.go.dev — builtin.max: https://pkg.go.dev/builtin#max
