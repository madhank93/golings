# If

`if` is the plainest thing in Go, and it is worth one chapter anyway — because
its *shape* is what most Go code looks like. There is no ternary operator, no
pattern matching, no truthiness. What there is instead is a statement form that
scopes a variable to the branch that needs it, and a convention (the early
return) that keeps the happy path flat down the left-hand side of the function.

Both exercises here are small: return the larger of two numbers, and map three
inputs to three outputs. The interesting part is what the idiomatic solution
looks like once the compiler's rules are taken seriously.

## 1. The condition is a `bool`, not a value

```go
if x { … }        // only if x is a bool
if x != 0 { … }   // what you write for a number
if s != "" { … }  // and for a string
```

No implicit conversion, so an integer, a string, or a nil pointer is never
"truthy". This removes a whole class of bugs (`if err` when you meant
`if err != nil`) at the cost of a few extra characters.

The parentheses are gone, but the braces are mandatory — even for a one-line
body. Combined with `gofmt`, that removes the dangling-else and
misleading-indentation families of bug entirely.

## 2. The statement form scopes the variable

```go
if v, err := strconv.Atoi(s); err == nil {
    use(v)
}
// v and err do not exist here
```

An `if` may run a short statement before its condition, separated by a
semicolon. Everything it declares lives only inside the `if` and its `else`
branches. This is the most common `if` in real Go — it keeps `err` from leaking
into the rest of the function, so the next call's `err` cannot be confused with
this one's.

## 3. Early return beats nesting

The Go style is to handle the exceptional case and leave, so the main path never
indents:

```go
// idiomatic
func process(s string) (int, error) {
    v, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    if v < 0 {
        return 0, errors.New("negative")
    }
    return v * 2, nil
}
```

Compare it with the `if/else` version, where the real work drifts right with
every check. `if2` is written in that nested shape on purpose; the direct fix is
an `else if` chain, and the version worth writing is a sequence of `if`s that
each return.

```ascii
nested                        early return
------                        ------------
if ok {                       if !ok {
    if ok2 {                      return err
        work()                }
    } else {                  if !ok2 {
        return err                return err
    }                         }
} else {                      work()
    return err
}
```

`else` after a block that always returns is redundant, and linters say so.

## 4. What Go does not have

- **No ternary.** `max := a > b ? a : b` does not exist; write the `if`, or use
  the built-in `max`/`min` (Go 1.21) for numbers, which is exactly `if1`'s job
  in one call.
- **No truthiness**, as above.
- **No `unless`, no statement modifiers.** One conditional form, plus `switch`
  for the multi-branch case — which the next chapter argues is the better tool
  the moment you have three branches.

## Gotchas

- **`:=` inside an `if` shadows** an outer variable of the same name. Usually
  what you want for `err`; occasionally the reason an assignment "did nothing".
- **A variable declared in the `if` statement is visible in `else`** too — the
  scope is the whole `if/else` chain, not just the first block.
- **Comparing floats with `==`** is a bug waiting for a rounding error; compare
  against a tolerance.
- **`if err != nil` is the check, not `if err`.** The latter does not compile,
  which is the point.
- **Empty branches** (`if cond {} else { … }`) mean the condition wants
  inverting.

## The exercises

- **if1** — return the larger of two numbers using only `if`.
- **if2** — map "fizz" → "foo", "fuzz" → "bar", anything else → "baz", and see
  how an `if/else` chain reads next to early returns.

## Source references

- [Go spec: If statements](https://go.dev/ref/spec#If_statements)
- [Effective Go: If](https://go.dev/doc/effective_go#if) — the early-return
  style, from the source
- [Go Code Review Comments: indent error flow](https://go.dev/wiki/CodeReviewComments#indent-error-flow)
- [pkg.go.dev: builtin.max](https://pkg.go.dev/builtin#max) (Go 1.21)
- [A Tour of Go: If](https://go.dev/tour/flowcontrol/5)

**Next: [switch](../switch/) →** — the same decision with three or more
branches, and Go's version of the fallthrough rules.
