## if2 — chaining conditions

```go
func fooIfFizz(fizzish string) string {
    if fizzish == "fizz" {
        return "foo"
    } else if fizzish == "fuzz" {
        return "bar"
    }
    return "baz"
}
```

**Why it works**

- Each condition is tested in order and the first match returns. The final
  `return` covers everything else, which is what the "neither" case needs.

**Common mistake**

- Comparing with a bare value — `if fizzish { … }`. Go has no truthiness: a
  condition must already be a `bool`, so the comparison is mandatory.

**Key detail:** three or more branches on one value is `switch`'s shape, and the
switch version reads better because the cases line up in a column:

```go
switch fizzish {
case "fizz":  return "foo"
case "fuzz":  return "bar"
default:      return "baz"
}
```

An `if` may also run a statement first — `if v, err := f(); err == nil` — which
scopes `v` and `err` to the branch and is the most common `if` in real Go.

**See also:** if1 (early return) · switch1 / switch3 (the same decision as a
switch) · the [chapter](../README.md)

**References**

- Go spec — If statements: https://go.dev/ref/spec#If_statements
- Effective Go — If: https://go.dev/doc/effective_go#if
