## slices2 — [low:high] excludes high

```go
func lastTwo(names [4]string) []string {
    return names[2:4] // index 2 and 3
}
```

**Why it works**

- Slicing takes `low` **inclusive** and `high` **exclusive**, so `[2:4]` yields
  two elements and `high-low` is always the length. `names[2:3]` would give one.

**Under the hood**

- Slicing an array produces a slice header pointing **into that array** — no
  copying happens. Writing `lastTwo(names)[0] = "X"` would change `names[2]` if
  the array were addressable, and the returned slice keeps the whole array alive
  as long as it exists.

**Common mistake**

- Reading `[2:4]` on a 4-element value as out of range. The high bound may equal
  the length (and in fact may go up to the *capacity*) — it names the position
  after the last element you want.

**Key detail:** both bounds are optional. `s[:n]` starts at 0, `s[n:]` runs to the
end, and `s[:]` is the whole thing — the usual way to get a slice over an array.

**See also:** slices1 (len vs cap) · slices4 (out-of-range bounds) ·
strings2 (slicing bytes, not characters) · the [chapter](../README.md)

**References**

- Go spec — Slice expressions: https://go.dev/ref/spec#Slice_expressions
- Go blog — Go Slices: usage and internals: https://go.dev/blog/slices-intro
