# Primitive types

Go's built-in types are deliberately boring: booleans, sized integers, two
floats, two complex types, and strings. What is not boring is the rule around
them — **Go never converts numeric types for you**. Adding an `int` to an
`int64` does not compile. That looks pedantic next to C or Python until the
first time it stops a silent truncation in code handling money or byte offsets.

These five exercises are all compile errors: a type that does not exist, a
variable used before it is declared, a byte literal with nothing in it. The
lesson underneath them is that in Go the type is part of the variable, chosen
once, and the compiler will not quietly bend it.

## 1. The type list

| Group | Types | Notes |
|---|---|---|
| Boolean | `bool` | `true` / `false`; no truthiness, no `0 == false` |
| Signed | `int8` `int16` `int32` `int64` `int` | `int` is 64-bit on modern platforms, but that is *implementation defined* |
| Unsigned | `uint8` `uint16` `uint32` `uint64` `uint` `uintptr` | wraps around on overflow rather than trapping |
| Float | `float32` `float64` | IEEE-754; `float64` is the default for an untyped float constant |
| Complex | `complex64` `complex128` | rarely used outside numerics |
| Text | `string` | immutable bytes, covered in its own chapter |

Two aliases matter constantly: **`byte` is `uint8`** and **`rune` is `int32`**.
They are the same types under different names, chosen to say what the value
means — a byte of data versus a Unicode code point.

There is no `integer` and no `float` (primitive_types5). Names come from the
table above, and a "big enough, machine-natural" integer is spelled `int`.

## 2. Conversion is explicit, always

```go
var i int = 42
var f float64 = float64(i)   // required — `= i` does not compile
var u uint8 = uint8(i)
```

The conversion syntax `T(v)` is the only way across. What it does *not* do is
check that the value fits:

```go
var big int = 300
var b byte = byte(big)   // compiles; b is 44 — the low 8 bits
```

Conversion truncates silently, so a range check is yours to write when the input
is not already known-good. Float → int truncates toward zero (`int(2.9)` is 2,
not 3); use `math.Round` first if you wanted rounding.

## 3. Untyped constants bend, variables do not

This compiles, and looks like it breaks the rule above:

```go
var f float64 = 3        // 3 is an untyped constant, becomes float64
var r rune = 'A'         // 'A' is an untyped rune constant
const big = 1 << 40      // no type; fits wherever it is used
```

An untyped constant has no type until context gives it one, and the compiler
evaluates it at arbitrary precision. The restriction only applies to *typed*
values — once something is stored in a variable, its type is fixed and
conversion becomes mandatory. That asymmetry is why `x * 2` is fine for any
numeric `x` while `x * y` across two different numeric types is not.

## 4. Bytes, characters, and `''`

```go
var b1 byte = 110   // the number 110
var b2 byte = 'n'   // the character 'n' — also 110
```

Single quotes make a **rune literal**: a single Unicode code point, whose value
is a number. `'n'` is `110`, `'A'` is `65`, `'世'` is `19990`. Empty quotes
(primitive_types4) hold no code point at all, so there is nothing for the
compiler to evaluate.

Double quotes make a string; backticks make a raw string that ignores escapes
and spans lines. A `byte` can hold `'n'` because it fits in 8 bits — `'世'`
would not, and needs a `rune`.

## 5. Formatting verbs, since every exercise here prints

`fmt.Printf` verbs are typed, and a mismatch is caught by `go vet` (which this
repo's lint step runs):

| Verb | For |
|---|---|
| `%d` | integers |
| `%s` | strings (and anything with a `String()` method) |
| `%f` / `%.2f` | floats, optionally with precision |
| `%t` | booleans |
| `%q` | quoted string — shows the empty string and whitespace clearly |
| `%v` / `%+v` | any value / with struct field names |
| `%T` | the value's *type*, the quickest way to answer "what is this?" |

## Gotchas

- **Sized types are for wire formats and memory layout**, not everyday code.
  Default to `int` for loop counters, lengths, and IDs; reach for `int64` when
  you need a guaranteed width.
- **Unsigned arithmetic wraps.** `var u uint8 = 0; u--` gives 255, with no
  panic. Loop counters that count down are safer as `int`.
- **`int` is not `int64`** even where both are 64 bits — passing one where the
  other is expected still needs a conversion.
- **Float equality is unreliable.** `0.1 + 0.2 != 0.3`. Compare with a
  tolerance, or use integers scaled by 100 for money.
- **Integer division truncates**: `3 / 2` is `1`. Convert first if you wanted
  `1.5`.
- **Overflow is silent in Go.** No panic, no trap — the value wraps.

## The exercises

- **primitive_types1** — reassign a `bool` between two checks.
- **primitive_types2** — declare the string before formatting it with `%s`.
- **primitive_types3** — same, for two variables at once.
- **primitive_types4** — put a character in the rune literal so the `byte` has
  a value.
- **primitive_types5** — `integer` and `float` are not Go type names; use the
  real ones.

## Source references

- [Go spec: Types](https://go.dev/ref/spec#Types) ·
  [Numeric types](https://go.dev/ref/spec#Numeric_types) ·
  [Conversions](https://go.dev/ref/spec#Conversions)
- [Go blog: Constants](https://go.dev/blog/constants) — the untyped-constant
  rules in full
- [pkg.go.dev: fmt](https://pkg.go.dev/fmt) — the verb table
- [A Tour of Go: Basic types](https://go.dev/tour/basics/11)

**Next: [if](../if/) →** — the first control flow, and the statement form that
keeps error variables tightly scoped.
