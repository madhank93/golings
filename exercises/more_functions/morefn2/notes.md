## morefn2 — variadic parameters and the spread

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)   // 6
sum(nums...)   // spread an existing slice
```

**Why it works**

- Inside the function `nums` is an ordinary `[]int`, so `range` sums it. At the
  call site the compiler collects the arguments into that slice — and `sum()`
  passes a `nil` slice, which ranges zero times and gives 0 with no special
  case.

**Under the hood**

- `slice...` at a call site passes the existing slice **without copying**: the
  function receives a header pointing at the same backing array, so writing to
  `nums[0]` mutates the caller's data. Copy with `slices.Clone` if the function
  keeps or modifies it.

**Common mistake**

- `sum(nums)` instead of `sum(nums...)` — passing one `[]int` where `int`s are
  expected. The compiler catches this one; the aliasing above it does not.

**Key detail:** the variadic parameter must be **last** and there can be only one.
This is how `fmt.Printf(format string, a ...any)` and `append(s, vals...)` are
declared — the mechanism is in the first line of Go anyone writes.

**See also:** morefn1 (recursion) · slices4 (`append` and backing arrays) ·
functions3 (exact arity otherwise) · the [chapter](../README.md)

**References**

- Go spec — Passing arguments to ... parameters: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
- pkg.go.dev — slices.Clone: https://pkg.go.dev/slices#Clone
