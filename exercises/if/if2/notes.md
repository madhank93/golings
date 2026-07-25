## if2 — chain conditions with else if

```go
if fizzish == "fizz" {
    return "foo"
} else if fizzish == "fuzz" {
    return "bar"
}
return "baz"
```

**Why it works**

- Each branch handles one input; the final `return "baz"` is the catch-all when
  neither condition matched.

**Nuance:** `else if` chains are checked **top to bottom** and stop at the first
match. The tests also show `if result := fooIfFizz(...); result != ... {` — an
**if with a short statement**: `result` is scoped to the `if`/`else` only.

**References**

- Go by Example — If/Else: https://gobyexample.com/if-else
