## modern2 — range over an integer

```go
func sumTo(n int) int {
    total := 0
    for i := range n { // i = 0, 1, ..., n-1
        total += i
    }
    return total
}
```

**Why it works**

- Go 1.22 lets `range` take an integer directly: `for i := range n` runs with
  `i` from 0 to `n-1`. No slice, no three-clause loop, no off-by-one to get
  wrong.

**Common mistake**

- Expecting it to include `n`. The bound is exclusive, like every other range in
  Go — `range 5` yields 0…4, which is why `sumTo(5)` is 10 and not 15.

**Key detail:** `for range n` without a variable is the "do this n times" loop,
and ranging 0 or a negative value runs zero times — no guard needed. The
three-clause form is still right when you need a different start, a step, or a
descending loop.

**See also:** modern1 · range1 (ranging collections) · morefn1 (loops vs
recursion) · the [chapter](../README.md)

**References**

- Go 1.22 release notes: https://go.dev/doc/go1.22#language
- Go spec — For statements with range clause: https://go.dev/ref/spec#For_range
