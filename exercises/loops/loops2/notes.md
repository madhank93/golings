## loops2: a while loop is a `for`

```go
func halvings(n int) int {
    steps := 0
    for n > 1 {
        n /= 2
        steps++
    }
    return steps
}
```

**Why it works**

- `for n > 1` checks the condition before every iteration, exactly like
  `while (n > 1)` elsewhere. Each pass halves `n`, so the condition is
  guaranteed to become false.

**Common mistake**

- Forgetting `n /= 2`. The condition never changes and the test hangs until
  golings times it out.

**Key detail:** changing `n` inside the function is safe. Go passes the
argument by value, so the caller's variable is untouched.

**See also:** loops1 · loops3 · the [chapter](../README.md)

**References**

- A Tour of Go, For is Go's "while": https://go.dev/tour/flowcontrol/3
