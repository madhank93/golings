## testing2: a check is an `if`

```go
func TestAbs(t *testing.T) {
    got := abs(-3)
    if got != 3 {
        t.Errorf("abs(-3) = %d, want 3", got)
    }
}
```

**Why it works**

- `t.Errorf` does not compare anything. It records a failure and a message.
  The `if` is the assertion, and Go has no built-in `assert` to replace it.

**Common mistake**

- Deleting the `t.Errorf` to get a green run. The test passes, but it no longer
  checks `abs` at all. Change `abs` to `return n` and confirm the test fails.

**Key detail:** the message format `abs(-3) = 3, want 3` (input, got, want)
is the Go convention. It reads like a sentence and tells you what to fix
without opening the file.

**See also:** testing1 · testing3 · the [chapter](../README.md)

**References**

- Go Code Review Comments, Useful Test Failures: https://go.dev/wiki/CodeReviewComments#useful-test-failures
