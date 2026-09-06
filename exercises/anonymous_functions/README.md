# Anonymous functions & closures

A function literal is a function written without a name, inline, where it is
needed. It has a type like any other function (`func(string)`), so it can be
called immediately, stored in a variable, passed as an argument, or returned.

What makes it more than shorthand is what happens when it mentions a variable
from the enclosing scope: the literal becomes a **closure**, and that variable
outlives the function it was declared in. That single mechanism is behind
`defer`, callbacks, middleware, iterators, and every "give me a function that
remembers something" pattern in Go.

## 1. Literals, and calling one on the spot

```go
func(name string) {
    fmt.Printf("Hello %s", name)
}("Gopher")        // the trailing (...) invokes it
```

The parentheses at the end are the call. Without an argument list that matches
the parameters, there is nothing to invoke — that is `anonymous_functions1`.
This immediately-invoked form is rare in Go; the common shapes are assigning the
literal, passing it, or `go func(){…}()`.

Assigning one means the variable's type is the **signature**:

```go
var sayBye func(name string)      // a variable of function type
sayBye = func(name string) {      // the literal must match exactly
    fmt.Printf("Bye %s", name)
}
```

`anonymous_functions2` is a mismatch — a `func()` assigned to a
`func(string)` variable, with a body referring to a parameter that was never
declared. Function types are structural: parameter types and result types, in
order. Names do not participate.

## 2. Closures capture variables, not values

```go
func updateStatus() func() string {
    var index int                        // lives on after updateStatus returns
    orderStatus := map[int]string{1: "TO DO", 2: "DOING", 3: "DONE"}

    return func() string {
        index++
        return orderStatus[index]
    }
}

next := updateStatus()
next()   // "TO DO"
next()   // "DOING"
```

`index` is declared in `updateStatus` and referenced by the returned literal, so
it cannot die when `updateStatus` returns. The compiler's escape analysis moves
it to the heap and the closure holds a reference — **not a copy**. Each call to
`updateStatus` produces a fresh `index`; two closures from two calls are
independent, and two closures from the *same* call share one.

```ascii
next := updateStatus()        other := updateStatus()

  next ──► closure ──► index(1)   other ──► closure ──► index(2)
                        ▲                                 ▲
        each call to updateStatus makes a new variable

  a := updateStatus()
  b := a                       both names, ONE closure, ONE index
```

That sharing is the whole point (a counter that advances) and the whole hazard
(two goroutines mutating one captured variable is a data race — guard it, as
`concurrent1` does).

## 3. Where closures actually show up

```go
defer func() { … }()                       // defer, every time
go func() { … }()                          // goroutines
sort.Slice(xs, func(i, j int) bool { … })  // callbacks
http.HandleFunc("/", func(w, r) { … })     // handlers
func(next http.Handler) http.Handler { … } // middleware
func(yield func(int) bool) { … }           // range-over-func iterators (1.23)
```

Two idioms are worth naming:

- **Functional options** — a constructor takes `...func(*Server)`, and each
  option is a closure over the value it will set. That is how Go does optional
  parameters.
- **Dependency injection by function** — pass `func(context.Context) (T, error)`
  instead of an interface with one method. Same substitutability, less
  ceremony (`di1`).

## 4. The loop-variable change

Before **Go 1.22**, all iterations of a loop shared one variable, so closures
created inside a loop all captured the same one:

```go
for _, v := range items {
    go func() { use(v) }()   // pre-1.22: every goroutine saw the LAST v
}
```

The fix used to be `v := v` at the top of the body. Since 1.22 each iteration
declares a fresh variable, so the plain closure captures what you expect. This
repo is on Go 1.26 — write it plainly, and recognise the old `v := v` line for
what it is when reading older code.

## Gotchas

- **Closures capture by reference.** Mutating the captured variable after the
  closure is created changes what the closure sees.
- **A captured variable escapes to the heap**, so a closure is not free — usually
  irrelevant, occasionally visible in a hot path (`go build -gcflags='-m'`).
- **Two closures from one call share state**; from two calls they do not.
- **Recursive literals need a declared variable first**: `var f func(int) int`
  then `f = func(n int) int { … f(n-1) … }`.
- **Function types must match exactly** — parameter and result types, in order.
- **`defer func(){…}()` runs at return; `defer f()` evaluates `f`'s arguments
  now.** See the `defer` chapter.

## The exercises

- **anonymous_functions1** — an inline literal needs an argument to be called
  with.
- **anonymous_functions2** — make the literal match the declared function type,
  and use its parameter.
- **anonymous_functions3** — return a closure that keeps `index` alive across
  calls.

## Source references

- [Go spec: Function literals](https://go.dev/ref/spec#Function_literals) ·
  [Function types](https://go.dev/ref/spec#Function_types)
- [Go blog: Fixing for loops in Go 1.22](https://go.dev/blog/loopvar-preview)
- [Effective Go: Function literals](https://go.dev/doc/effective_go#func_literals)
- [Go by Example: Closures](https://gobyexample.com/closures)

**Next: [defer](../defer/) →** — the closure Go runs for you on the way out.
