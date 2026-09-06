# Defer

`defer` schedules a call to run when the surrounding **function** returns —
not when the block ends, and not at some point later. It runs on every exit
path: a normal return, an early return, a `t.Fatal`, even a panic on its way up
the stack.

That last property is what makes it the cleanup mechanism in Go. There is no
`finally`, no RAII, no destructors; there is `defer`, sitting one line below the
thing it cleans up, where a reader can see the pair together.

## 1. LIFO, and when the arguments are evaluated

```go
func order() (seq []int) {
    defer func() { seq = append(seq, 1) }()
    defer func() { seq = append(seq, 2) }()
    defer func() { seq = append(seq, 3) }()
    return   // runs 3, then 2, then 1  →  [3 2 1]
}
```

Deferred calls are pushed onto a stack and run **last in, first out** — that is
`defer1`. LIFO is the right order for cleanup: the last resource acquired is the
first released, so a lock taken inside an open file is released before the file
closes.

The other half of the rule catches everyone once:

```go
i := 0
defer fmt.Println(i)   // prints 0 — i is evaluated NOW
i = 42
```

**Arguments are evaluated when the `defer` statement runs; the call happens
later.** Wrap the work in a literal — `defer func(){ fmt.Println(i) }()` — when
you want the value as it is at return time.

## 2. Guaranteed cleanup

```go
func process(r *Resource, early bool) {
    defer r.Close()   // top of the function, right after acquiring

    if early {
        return        // Close still runs
    }
    // … more work, more returns …
}
```

`defer2` is this shape, and it is the shape of most real Go:

```go
f, err := os.Open(name)
if err != nil {
    return err
}
defer f.Close()

mu.Lock()
defer mu.Unlock()

resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
```

The pattern is always: acquire, check the error, `defer` the release on the next
line. Deferring **before** the error check is the bug — you would call `Close`
on a `nil` file.

## 3. Named results: `defer` can change what is returned

A `return x` in a function with named results assigns `x` to the result variable
and *then* runs the defers, so a deferred closure can still modify it:

```ascii
return x
   │
   ├─ 1. assign x to the named result
   ├─ 2. run deferred calls  (they can still change it)
   └─ 3. actually return
```

```go
func safeRun(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)   // sets the RESULT
        }
    }()
    fn()
    return nil
}
```

This only works with a **named** result — an anonymous `(error)` gives the
closure nothing to assign to. It is the mechanism behind the recover pattern in
`errors4`, and behind `defer func(){ err = errors.Join(err, f.Close()) }()`,
which is how you stop a failing `Close` from being silently dropped.

## 4. What it costs, and where it does not run

Since Go 1.14 most defers are **open-coded** — the compiler inlines them at each
return path, so a `defer` in a straight-line function costs a few nanoseconds.
Ones inside a loop, or under a condition the compiler cannot resolve, fall back
to a heap-allocated defer record and cost more.

Two situations where a deferred call never happens:

- **`os.Exit`** terminates immediately. So does a signal-killed process.
- **A goroutine that never returns** — the defers are attached to the function
  call, so a blocked goroutine's cleanup never runs either.

And one where the timing surprises people: **`defer` in a loop** stacks up until
the *function* returns, not the iteration. Opening a thousand files in a loop
with `defer f.Close()` holds a thousand descriptors. Put the body in its own
function (or a closure called per iteration) so the defer fires each time.

## Gotchas

- **Arguments evaluate immediately**; only the call is delayed.
- **`defer` in a loop accumulates** — one function, one stack of defers.
- **Deferring before checking the error** defers a call on a nil value.
- **`defer mu.Unlock()` immediately after `Lock()`** is the idiom; anything else
  invites an unlock that never happens.
- **A deferred `Close` on a writer swallows its error.** Assign it to a named
  result if the write must be durable.
- **Only a named result can be modified** by a deferred closure.
- **`recover()` only works when called directly by a deferred function** — not
  nested one level deeper.

## The exercises

- **defer1** — three deferred appends run in reverse, shaping a named result.
- **defer2** — defer the cleanup at the top so it runs on the early-return path
  too.

## Source references

- [Go spec: Defer statements](https://go.dev/ref/spec#Defer_statements) ·
  [Return statements](https://go.dev/ref/spec#Return_statements) — the
  assign-then-defer ordering
- [Effective Go: Defer](https://go.dev/doc/effective_go#defer)
- [Go blog: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Go 1.14 release notes](https://go.dev/doc/go1.14#runtime) — open-coded defers

**Next: [errors](../errors/) →** — the values those cleanups protect, and the
chain that carries them up.
