## enums1 — iota numbers the block

```go
type Weekday int

const (
    Sunday Weekday = iota // 0
    Monday                // 1 — the expression repeats
    Tuesday               // 2
    Wednesday
    Thursday
    Friday
    Saturday              // 6
)
```

**Why it works**

- A `const` line with no `= …` **repeats the previous expression**, and `iota` is
  the line's index within the block. So one `= iota` at the top numbers every
  following name automatically.

**Under the hood**

- `iota` counts **ConstSpec lines**, not constants, and resets to 0 in every
  `const` block — not per file or package. That is why `A, B = iota, iota` gives
  both the value 0, and why shifting tricks work:
  `KB = 1 << (10 * iota)` on successive lines yields KB, MB, GB.

**Common mistake**

- Reordering or inserting a name later. Every value after the insertion point
  silently renumbers — harmless in memory, corrupting if those numbers were
  written to a database or a wire format. Assign explicit numbers for values
  that persist.

**Key detail:** `type Weekday int` is what makes this more than named ints — a
plain `int` variable will not pass where a `Weekday` is expected. Untyped
constants still convert freely, though, so `schedule(3)` compiles; the compiler
does no range check and `Weekday(42)` is a valid value.

**See also:** enums2 (`String()`) · typealias1 (why the named type is distinct) ·
switch3 · the [chapter](../README.md)

**References**

- Go spec — Iota: https://go.dev/ref/spec#Iota
- Effective Go — Constants: https://go.dev/doc/effective_go#constants
