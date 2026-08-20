## anonymous_functions3 — a closure remembers

```go
func updateStatus() func() string {
    var index int
    orderStatus := map[int]string{1: "TO DO", 2: "DOING", 3: "DONE"}

    return func() string {
        index++
        return orderStatus[index]
    }
}
```

**Why it works**

- The returned literal references `index` and `orderStatus` from the enclosing
  function, so both stay alive after `updateStatus` returns. Each call advances
  `index` and returns the next status.

**Under the hood**

- A closure captures **variables, not values** — it holds a reference. Escape
  analysis moves `index` to the heap because its address outlives the call
  (`go build -gcflags='-m'` prints `moved to heap: index`). Every call to
  `updateStatus` creates a *fresh* `index`, so two closures from two calls are
  independent; two names for the same closure share one.

**Common mistake**

- Sharing a captured variable across goroutines without synchronisation. It is
  one variable, so concurrent writes are a data race — `go test -race` flags it.

**Key detail:** this pattern is how Go does stateful callbacks: iterators,
rate limiters, memoisers, and the functional-options pattern
(`...func(*Server)`) are all closures over their configuration.

**See also:** anonymous_functions2 · concurrent1 (closures in goroutines) ·
iter1 (closures as iterators) · the [chapter](../README.md)

**References**

- Go spec — Function literals: https://go.dev/ref/spec#Function_literals
- A Tour of Go — Closures: https://go.dev/tour/moretypes/25
