## modern5 — the loop must declare the variable

```go
for _, item := range items { // := lets the loop declare it
    fns = append(fns, func() string { return item })
}
```

**Why it works**

- Since Go 1.22 each iteration of a `for` loop declares **new** loop variables, so
  every closure captures its own `item`. The broken version hoists `item` above
  the loop and merely assigns it with `=` — one variable shared by every closure,
  all reporting the last value.

**Under the hood**

- The 1.22 change applies to variables the **loop declares**, not to any variable
  the loop touches. Assigning to an outer variable is still assignment, so the
  old pre-1.22 behaviour is still reachable — which is what this exercise
  demonstrates on purpose.

**Common mistake**

- Adding `item := item` at the top of the body out of habit. It was the standard
  workaround before 1.22 and is now redundant; linters flag the copy. Recognise
  it in older code rather than writing it.

**Key detail:** the same rule governs pointers, not just closures.
`&item` inside a pre-1.22-shaped loop yields the same address every iteration —
so a slice of pointers ends up with every element aimed at one variable.

**Note on modules:** the semantics are gated by the `go` line in `go.mod`. A
module declaring `go 1.21` keeps the old behaviour even on a new toolchain.

**See also:** anonymous_functions3 (closures capture variables) ·
concurrent1 (goroutines in a loop) · range3 · the [chapter](../README.md)

**References**

- Go blog — Fixing for loops in Go 1.22: https://go.dev/blog/loopvar-preview
- Go 1.22 release notes: https://go.dev/doc/go1.22#language
