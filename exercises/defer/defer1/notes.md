## defer1 — defers run LIFO

```go
func order() (seq []int) {
    defer func() { seq = append(seq, 1) }()
    defer func() { seq = append(seq, 2) }()
    defer func() { seq = append(seq, 3) }()
    return // runs 3, 2, 1 -> [3 2 1]
}
```

**Why it works**

- Deferred calls are pushed on a stack and run **last in, first out** when the
  function returns. The last `defer` written is the first to run, so appending
  1, 2, 3 produces `[3 2 1]`.

**Under the hood**

- `return` in a function with **named results** is three steps: assign the result
  variable, run the deferred calls, then actually return. That ordering is why
  these closures can still append to `seq` — and why a deferred closure can
  change what the caller receives.

**Common mistake**

- Expecting `defer fmt.Println(i)` to print `i`'s final value. Arguments are
  evaluated **when the `defer` statement runs**; only the call is delayed. Wrap
  the work in a literal to read the variable at return time.

**Key detail:** LIFO is the correct order for cleanup — the last resource
acquired is released first, so a lock taken inside an open file is unlocked
before the file closes.

**See also:** defer2 (guaranteed cleanup) · errors4 (a deferred closure setting
a named result) · anonymous_functions3 · the [chapter](../README.md)

**References**

- Go spec — Defer statements: https://go.dev/ref/spec#Defer_statements
- Go blog — Defer, Panic, and Recover: https://go.dev/blog/defer-panic-and-recover
