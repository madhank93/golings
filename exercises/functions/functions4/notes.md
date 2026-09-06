## functions4 — declare what the function returns

```go
func addNumbers(a int, b int) int {
    return a + b
}
```

**Why it works**

- The broken version has `return a + b` in a function whose signature promises
  nothing, so the value has nowhere to go: `too many return values`. Naming the
  result type `int` completes the contract.

**Common mistake**

- Forgetting that **every path** must return. A function with a result type and
  a `return` only inside an `if` fails with `missing return`, even when you can
  see the loop always exits — the compiler will not reason about that for you.

**Key detail:** results can be plural, and usually are: `func divide(a, b int)
(int, error)` is the signature Go writes instead of throwing. Results may also
be named — `func split(sum int) (x, y int)` — which documents what the two
`int`s mean and lets a `defer` adjust them before they leave.

**See also:** functions2 · errors1 (the `(T, error)` pair) · defer1 (a `defer`
that shapes a named result) · the [chapter](../README.md)

**References**

- Go spec — Return statements: https://go.dev/ref/spec#Return_statements
- Effective Go — Multiple return values: https://go.dev/doc/effective_go#multiple-returns
