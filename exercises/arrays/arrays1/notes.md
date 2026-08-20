## arrays1 — indexes start at zero

```go
func first(colors [3]string) string { return colors[0] }
func last(colors [3]string) string  { return colors[2] }
```

**Why it works**

- A `[3]string` has valid indexes 0, 1 and 2. The first element is `[0]` and the
  last is `len-1`, which is `[2]` — not `[3]`.

**Common mistake**

- Reaching for `colors[3]` as "the third". Go catches a constant index at
  **compile time** (`invalid argument: index 3 out of bounds [0:3]`); a computed
  index that goes out of range panics at run time instead. Either way there is
  no silent read past the end.

**Key detail:** the length is part of the type — `[3]string` and `[4]string` are
different types, so `len(colors)` is a compile-time constant here. A function
that should accept any number of strings takes a slice, `[]string`.

**See also:** arrays2 (one element type) · slices2 (`[low:high]` bounds) ·
slices4 (the panics) · the [chapter](../README.md)

**References**

- Go spec — Array types: https://go.dev/ref/spec#Array_types
- A Tour of Go — Arrays: https://go.dev/tour/moretypes/6
