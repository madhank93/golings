## loops1: an off-by-one condition

```go
func sumTo(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}
```

**Why it works**

- The loop has to run for every `i` from 1 up to and including `n`. With
  `i < n` it stops after `n-1`.

**Common mistake**

- Starting at `i := 0` and keeping `i < n`. That also runs `n` times, but adds
  0 to `n-1`, so `sumTo(4)` is still 6.

**Key detail:** read a loop by its edges. What is `i` on the first iteration,
and what is it on the last? If both match what you want, the middle usually
takes care of itself.

**See also:** loops2 · range1 · the [chapter](../README.md)

**References**

- A Tour of Go, For: https://go.dev/tour/flowcontrol/1
