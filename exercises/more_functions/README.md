# More functions

Two techniques that come up constantly once the basics are in place: a function
that calls itself, and a function that accepts any number of arguments. Both are
small additions to the syntax from the previous chapter, and both have one sharp
edge worth knowing before you use them — an unbounded stack in the first case,
a shared backing array in the second.

## 1. Recursion, and where it stops

```go
func factorial(n int) int {
    if n <= 1 {
        return 1        // base case — the only way out
    }
    return n * factorial(n-1)
}
```

The base case is not a detail, it is the whole design: without it the calls
never stop returning to themselves and the program dies with
`runtime: goroutine stack exceeds 1000000000-byte limit` followed by
`fatal error: stack overflow`.

Go's stacks make this less dangerous than it sounds. A goroutine starts with a
2 KB stack and the runtime grows it by allocating a bigger one and copying —
so recursion depth is bounded by a limit (1 GB on 64-bit by default), not by a
small fixed frame count. Deep recursion is possible; it is just rarely the
clearest way to write a loop.

What Go does **not** do is guarantee tail-call elimination. `return f(n-1)` in
tail position still consumes a frame, so a recursive function processing a
million-element list will use a million frames. Write the loop for those; keep
recursion for genuinely recursive shapes — trees, directory walks, parsers.

```ascii
factorial(4)
  4 * factorial(3)
        3 * factorial(2)
              2 * factorial(1)
                    1          <- base case, unwinding starts
              2 * 1  = 2
        3 * 2  = 6
  4 * 6  = 24
```

## 2. Variadic parameters

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum()             // 0    — nums is an empty (nil) slice
sum(1, 2, 3)      // 6
```

Inside the function, `nums` is a plain `[]int`. The compiler builds that slice
at the call site from whatever arguments were passed, so `sum()` receives a
`nil` slice — `len` 0, `range` runs zero times, no special case needed.

The variadic parameter must be **last**, and there can be only one:

```go
func logf(prefix string, args ...any)   // fine
func bad(args ...any, suffix string)    // compile error
```

This is how `fmt.Printf(format string, a ...any)` and `append(s, vals...)` are
declared — variadic functions are not an exotic corner, they are in the first
line of code you ever wrote in Go.

## 3. Spreading a slice: `slice...`

```go
nums := []int{4, 5, 6}
sum(nums...)      // 15 — pass the existing slice as the arguments
sum(nums)         // compile error: []int is not int
```

The `...` suffix at a *call* site says "these are the arguments", the mirror of
the `...` in the declaration. Without it you are passing one `[]int` where
`int`s are expected.

One thing worth knowing: spreading does **not** copy. The function receives a
slice header pointing at the same backing array, so a variadic function that
writes to `nums[0]` mutates the caller's slice. Functions that keep or modify
their variadic argument should copy it first (`slices.Clone`).

## Gotchas

- **No base case, or a base case the arguments never reach** (`n-2` on odd
  input) is an infinite recursion — a stack overflow, not a hang.
- **Recursion has no tail-call optimisation in Go.** Depth costs frames.
- **`sum(nums)` vs `sum(nums...)`** — one is a type error, the other is the
  spread. The compiler catches this one, at least.
- **A spread slice is shared, not copied.** Mutating a variadic parameter
  mutates the caller's data.
- **`nums` is `nil`, not empty, when no arguments are passed.** `len(nums) == 0`
  either way, so `range` is safe — but `nums == nil` is `true`.
- **A variadic parameter is not optional arguments.** `f(a int, b ...int)` lets
  callers omit `b`, but every value in `b` must be the same type.

## The exercises

- **morefn1** — write `factorial` recursively, with `0! == 1` as the base case.
- **morefn2** — sum a variadic `...int`, then call it by spreading a slice.

## Source references

- [Go spec: Function types](https://go.dev/ref/spec#Function_types) (variadic
  parameters) · [Passing arguments to ... parameters](https://go.dev/ref/spec#Passing_arguments_to_..._parameters)
- [pkg.go.dev: fmt.Printf](https://pkg.go.dev/fmt#Printf) — the variadic
  signature you use every day
- [pkg.go.dev: slices.Clone](https://pkg.go.dev/slices#Clone) — when a variadic
  argument must be kept
- [Go by Example: Recursion](https://gobyexample.com/recursion) ·
  [Variadic Functions](https://gobyexample.com/variadic-functions)

**Next: [strings](../strings/) →** — the type every one of these exercises has
been printing, and why its length is not what you think.
