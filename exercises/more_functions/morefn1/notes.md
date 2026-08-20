## morefn1 — recursion needs a base case

```go
func factorial(n int) int {
    if n <= 1 {
        return 1 // base case: stops the recursion
    }
    return n * factorial(n-1)
}
```

**Why it works**

- Each call reduces `n` by one until it reaches the base case, which returns
  without recursing. The stack then unwinds, multiplying on the way back out:
  `4 * (3 * (2 * 1))`.

**Under the hood**

- Every call gets its own frame holding its own `n`. A goroutine starts with a
  2 KB stack and the runtime grows it by allocating a larger one and copying, so
  depth is bounded by a limit (1 GB on 64-bit), not a small fixed count. Go does
  **not** do tail-call elimination: `return f(n-1)` still costs a frame.

**Common mistake**

- A base case the argument never hits — recursing on `n-2` from an odd `n`, or
  forgetting `n <= 1` and testing only `n == 1`. The result is not a hang but
  `fatal error: stack overflow` after the stack limit is reached.

**Key detail:** use recursion for genuinely recursive shapes — trees, directory
walks, parsers. A linear sequence is clearer and cheaper as a loop.

**See also:** morefn2 (variadic) · functions4 (return types) ·
the [chapter](../README.md)

**References**

- Go spec — Function declarations: https://go.dev/ref/spec#Function_declarations
- Go by Example — Recursion: https://gobyexample.com/recursion
