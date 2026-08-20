## range3 — filter with range

```go
evenNumbers := []int{}
for _, n := range numbers {
    if n%2 == 0 {
        evenNumbers = append(evenNumbers, n)
    }
}
```

**Why it works**

- Range every value, test it, and `append` the ones that pass to a second slice.
  The result must be assigned back — `append` returns the new header.

**Under the hood**

- `range` evaluates its subject **once**, before the first iteration, so the loop
  walks `numbers` as it was at the start. Appending to a *different* slice inside
  the body (as here) is fine; appending to the one being ranged would not extend
  the loop.

**Common mistake**

- Mutating the loop variable and expecting the source to change. `n` is a **copy**
  of the element — `n = 0` changes nothing. Write through the index
  (`numbers[i] = 0`) when you mean to modify in place.

**Key detail:** since Go 1.22 each iteration gets a **fresh** loop variable, so
capturing `n` in a closure or goroutine now captures that iteration's value. The
old `n := n` shadow line is no longer needed; older code and older blog posts
predate the change.

**See also:** range1 · slices3 (`append`) · concurrent1 (loop variables in
goroutines) · the [chapter](../README.md)

**References**

- Go spec — For statements with range clause: https://go.dev/ref/spec#For_range
- Go blog — Fixing for loops in Go 1.22: https://go.dev/blog/loopvar-preview
