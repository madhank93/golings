## morefn1 — recursion needs a base case

```go
func factorial(n int) int {
    if n <= 1 {
        return 1 // base case: stops the recursion
    }
    return n * factorial(n-1) // recursive case: moves toward the base
}
```

**Why it works**

- `factorial` calls itself with a smaller `n` each time. The `n <= 1` **base
  case** ends the chain; without it the calls would never stop.

**Nuance:** every recursion needs (1) a base case that returns without recursing
and (2) a recursive step that provably moves **toward** it. Miss either and you
get infinite recursion → a stack-overflow crash.

**References**

- Go by Example — Recursion: https://gobyexample.com/recursion
