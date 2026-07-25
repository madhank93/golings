## modern1 — built-in min, max, clear

```go
min(a, b, c) // smallest
max(a, b, c) // largest
clear(m)     // empty a map (or zero a slice)
```

**Why it works**

- Go 1.21 added `min`, `max`, and `clear` as **built-ins** — no import, and `min`/
  `max` are variadic and generic over ordered types. `clear(m)` deletes every
  entry from a map in place.

**Nuance:** these are language built-ins, not functions in a package, so they
work on any ordered type without generics boilerplate. `clear` on a **map**
removes all keys; on a **slice** it zeroes every element (keeping the length).

**References**

- Go 1.21 release notes: https://go.dev/doc/go1.21#language
