## modern1 — built-in min, max, clear

```go
func bounds(a, b, c int) (lo, hi int) {
    return min(a, b, c), max(a, b, c)
}

func wipe(m map[string]int) { clear(m) }
```

**Why it works**

- Go 1.21 added `min`, `max` and `clear` as **built-ins** — no import, and they
  work for any ordered type (including strings) with two or more arguments.

**Under the hood**

- `clear` has two behaviours worth keeping straight. On a **map** it deletes
  every entry, leaving a usable empty map. On a **slice** it zeroes each element
  and leaves the **length unchanged** — to empty a slice you still write
  `s = s[:0]`.

**Common mistake**

- Still shipping a hand-rolled `func maxInt(a, b int) int`. Every codebase had
  one (often several, one per type); the built-ins replace all of them and are
  generic without the ceremony.

**Key detail:** these are language built-ins like `len` and `append`, not
functions in a package — so they cannot be shadowed by an import, but they *can*
be shadowed by a local variable named `min`, which then silently stops compiling
as a call.

**See also:** modern2 · maps1 (map basics) · slices1 · generics2 (`cmp.Ordered`,
the same idea as a constraint) · the [chapter](../README.md)

**References**

- Go 1.21 release notes: https://go.dev/doc/go1.21#language
- pkg.go.dev — builtin: https://pkg.go.dev/builtin#min
