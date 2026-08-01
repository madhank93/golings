## anonymous_functions3 — a closure that remembers state

```go
func updateStatus() func() string {
    index := 1
    orderStatus := map[int]string{1: "TO DO", 2: "DOING", 3: "DONE"}
    return func() string {
        index++
        return orderStatus[index]
    }
}
```

**Why it works**

- The returned function **closes over** `index` and `orderStatus`. Each call
  increments the *same* `index`, so successive calls return "DOING" then "DONE".

**Key detail:** the captured variables live on **past** `updateStatus`'s return —
they're promoted to the heap and shared by every call to the returned closure.
That persistent, private state is exactly what makes closures useful (counters,
generators, memoization).

**References**

- A Tour of Go — Closures: https://go.dev/tour/moretypes/25
- Go by Example — Closures: https://gobyexample.com/closures
